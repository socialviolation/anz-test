package main

import (
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

}
