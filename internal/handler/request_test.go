package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = gin.New()
	router.GET("/nthPrime/:n", GetPrime)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestGetPrime(t *testing.T) {
	req, err := http.NewRequest("GET", "/nthPrime/19", nil)
	if err != nil {
		t.Fatalf("Failed to create new HTTP GET request: %v", err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseBody response
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := response{Prime: 71}
	assert.Equal(t, expectedResponse, responseBody)
}

func TestParamNotIntegerParsable(t *testing.T) {
	req, err := http.NewRequest("GET", "/nthPrime/foo", nil)
	if err != nil {
		t.Fatalf("Failed to create new HTTP GET request: %v", err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestParamNegative(t *testing.T) {
	req, err := http.NewRequest("GET", "/nthPrime/-1", nil)
	if err != nil {
		t.Fatalf("Failed to create new HTTP GET request: %v", err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
