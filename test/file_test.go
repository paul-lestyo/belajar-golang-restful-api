package test

import (
	"github.com/paul-lestyo/belajar-golang-restful-api/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
