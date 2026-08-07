package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilityPackage(t *testing.T) {

	value := "redira"

	assert.NotEmpty(
		t,
		value,
	)
}
