package links

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortCodeGeneration(t *testing.T) {

	shortCode := "abc123"

	assert.NotEmpty(
		t,
		shortCode,
	)

	assert.Equal(
		t,
		6,
		len(shortCode),
	)
}
