package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {

	os.Setenv(
		"APP_NAME",
		"Redira",
	)

	defer os.Unsetenv(
		"APP_NAME",
	)

	assert.Equal(
		t,
		"Redira",
		os.Getenv("APP_NAME"),
	)
}

func TestEnvironmentVariableMissing(t *testing.T) {

	os.Unsetenv(
		"DATABASE_URL",
	)

	assert.Empty(
		t,
		os.Getenv("DATABASE_URL"),
	)
}
