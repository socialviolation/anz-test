package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	// Mainly want to check the port hasn't been changed
	assert.Equal(t, appName, "myapplication", "app name should be equal")
	assert.Equal(t, description, "pre-interview technical test", "description should be equal")
	assert.Equal(t, port, ":8080", "port should be 8080")
}

func TestGetVersion(t *testing.T) {
	os.Setenv("APP_VERSION", "0.1.2-test")
	os.Setenv("SHA", "949419ab387a681db0b447452090668cac8b6b55")

	result := getVersion()

	assert.Equal(t, len(result.App), 1)

	v := result.App[0]
	assert.Equal(t, v.Version, "0.1.2-test", "version should be os env var")
	assert.Equal(t, v.SHA, "949419ab387a681db0b447452090668cac8b6b55", "sha should be os env var")
	assert.Equal(t, v.Description, description, "description should be const value")
}

func TestVersionHandler(t *testing.T) {
	os.Setenv("APP_VERSION", "0.1.2-test")
	os.Setenv("SHA", "949419ab387a681db0b447452090668cac8b6b55")

	response := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/version", nil)

	setupRoutes().ServeHTTP(response, req)

	assert.Nil(t, err, "request error should be nil")
	assert.Equal(t, response.Code, http.StatusOK, "expected response to be 200")
	expected := `{"myapplication":[{"version":"0.1.2-test","lastcommitsha":"949419ab387a681db0b447452090668cac8b6b55","description":"pre-interview technical test"}]}`
	actual := strings.Trim(response.Body.String(), "\n")
	assert.Equal(t, actual, expected, "expected body to have version map")
}
