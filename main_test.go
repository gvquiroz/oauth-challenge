package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestHealthCheckRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestDefaultRouteWithJsonContentType(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
	assert.Equal(t, strings.Contains("application/json; charset=utf-8",w.Header().Get("ContentType")), true, "Content type must be application json")
}

func TestCountRouteWithJsonContentType(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/count", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// return empty json
	assert.Equal(t, "{}", w.Body.String())
	assert.Equal(t, strings.Contains("application/json; charset=utf-8",w.Header().Get("ContentType")), true, "Content type must be application json")
}

func TestCountWithAcents(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/count?input=papá", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// return empty json
	assert.Equal(t, `{"a":2,"p":2}` , w.Body.String())
	assert.Equal(t, strings.Contains("application/json; charset=utf-8",w.Header().Get("ContentType")), true, "Content type must be application json")
}

func TestCountWithDieresis(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/count?input=päpa", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// return empty json
	assert.Equal(t, `{"a":2,"p":2}` , w.Body.String())
	assert.Equal(t, strings.Contains("application/json; charset=utf-8",w.Header().Get("ContentType")), true, "Content type must be application json")
}

func TestCountWithSymbols(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/count?input=päpa!!", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// return empty json
	assert.Equal(t, `{"a":2,"p":2}` , w.Body.String())
	assert.Equal(t, strings.Contains("application/json; charset=utf-8",w.Header().Get("ContentType")), true, "Content type must be application json")
}