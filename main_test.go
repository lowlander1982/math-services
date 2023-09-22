package main

import (
	"math-services/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/pi", handlers.GetDigitsOfPi)
	return r
}

func TestDefaultDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "\"realPi\":\"3\"")
}

func TestTwoDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=2", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "\"realPi\":\"3.14\"")
}

func TestInvalidDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=abc", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "El valor de 'digits' debe ser un número no negativo")
}

func TestNegativeDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=-5", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "El valor de 'digits' debe ser un número no negativo")
}

func TestLargeDigits(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/pi?digits=1000", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "\"realPi\":\"3.141592653589793\"")
}
