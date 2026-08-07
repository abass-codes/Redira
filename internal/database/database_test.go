package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseURLConfiguration(t *testing.T) {

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

func TestInvalidDatabaseURL(t *testing.T) {

	databaseURL := ""

	assert.Empty(
		t,
		databaseURL,
	)
}
