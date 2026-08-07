package redirect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirectCodeValidation(t *testing.T) {

	code := "abc123"

	assert.NotEmpty(
		t,
		code,
	)

	assert.Equal(
		t,
		6,
		len(code),
	)
}
