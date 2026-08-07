package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectionConfiguration(t *testing.T) {

	databaseURL := "postgres://redira:password@localhost:5432/redira"

	assert.NotEmpty(
		t,
		databaseURL,
	)

	assert.Contains(
		t,
		databaseURL,
		"postgres",
	)
}
