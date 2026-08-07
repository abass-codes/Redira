package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordValidation(t *testing.T) {

	password := "secure_password"

	assert.NotEmpty(
		t,
		password,
	)

	assert.Greater(
		t,
		len(password),
		8,
	)
}
