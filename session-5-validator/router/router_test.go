// router/router_test.go
package router_test

import (
	"golang-advance/session-5-validator/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestPublicRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUserHandler := &MockUserHandler{}
	router.SetupRouter(r, mockUserHandler)

	t.Run("GET /users/:id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		require.Contains(t, w.Body.String(), "user found")
	})

	t.Run("GET /users", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		require.Contains(t, w.Body.String(), "all users")
	})
}

func TestPrivateRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUserHandler := &MockUserHandler{}
	router.SetupRouter(r, mockUserHandler)

	addAuth := func(req *http.Request) {
		req.SetBasicAuth("user", "pass")
	}

	t.Run("POST /users", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", nil)
		addAuth(req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusCreated, w.Code)
		require.Contains(t, w.Body.String(), "user created")
	})

	t.Run("PUT /users/:id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/users/1", nil)
		addAuth(req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		require.Contains(t, w.Body.String(), "user updated")
	})

	t.Run("DELETE /users/:id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		addAuth(req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		require.Contains(t, w.Body.String(), "user deleted")
	})
}

func TestPrivateRoutesUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUserHandler := &MockUserHandler{}
	router.SetupRouter(r, mockUserHandler)

	t.Run("POST /users - Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
		require.Contains(t, w.Body.String(), "Authorization basic token required")
	})

	t.Run("PUT /users/:id - Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
		require.Contains(t, w.Body.String(), "Authorization basic token required")
	})

	t.Run("DELETE /users/:id - Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
		require.Contains(t, w.Body.String(), "Authorization basic token required")
	})
}
