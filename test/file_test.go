package test

import (
	"testing"
	"vick1208/restful-api/simple"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	conn, cleanup := simple.InitializedConnect("Database")

	assert.NotNil(t, conn)
	cleanup()
}
