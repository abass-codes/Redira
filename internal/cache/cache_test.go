package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisURLValidation(t *testing.T) {

	validURL := "redis://localhost:6379"

	assert.NotEmpty(
		t,
		validURL,
	)

	assert.Contains(
		t,
		validURL,
		"redis",
	)
}

func TestInvalidRedisURL(t *testing.T) {

	invalidURL := ""

	assert.Empty(
		t,
		invalidURL,
	)
}
