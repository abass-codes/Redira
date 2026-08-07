package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerCreation(t *testing.T) {

	logger := New()

	assert.NotNil(
		t,
		logger,
	)

	logger.Sync()
}
