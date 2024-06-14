package router_test

import (
	"bytes"
	"encoding/json"
	"golang-advance/session-3-unit-test/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter_RootHandler(t *testing.T) {
	r := gin.Default()
	router.SetupRouter(r)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Halo dari Gin!"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestPostHandler_PositiveCase(t *testing.T) {
	r := gin.Default()
	router.SetupRouter(r)

	requestBody := map[string]string{"message": "Test message"}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/private/post", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid-token")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Test message"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestPostHandler_NegativeCase_BadRequest(t *testing.T) {
	r := gin.Default()
	router.SetupRouter(r)

	req, _ := http.NewRequest("POST", "/private/post", bytes.NewBufferString("{Invalid JSON}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid-token")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.Contains(t, w.Body.String(), "invalid character")
}

func TestPostHandler_NegativeCase_NoAuthHeader(t *testing.T) {
	r := gin.Default()
	router.SetupRouter(r)

	req, _ := http.NewRequest("POST", "/private/post", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	assert.Contains(t, w.Body.String(), "{\"error\":\"Authorization token required\"}")
}
