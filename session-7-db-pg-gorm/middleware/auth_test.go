package middleware_test

import (
	"golang-advance/session-7-db-pg-gorm/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		username       string
		password       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid credentials",
			username:       "user",
			password:       "pass",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "Invalid credentials",
			username:       "user",
			password:       "wrongpass",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid authorization token"}`,
		},
		{
			name:           "No credentials",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Authorization basic token required"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.Use(middleware.AuthMiddleware())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "OK")
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.username != "" || tt.password != "" {
				req.SetBasicAuth(tt.username, tt.password)
			}

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			require.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				require.Equal(t, tt.expectedBody, w.Body.String())
			} else {
				require.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
