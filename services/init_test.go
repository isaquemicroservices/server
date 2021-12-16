package services_test

import (
	"testing"

	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/stretchr/testify/assert"
)

func TestServices(t *testing.T) {
	// TestInitializeConnections function to test initializing connections with services gRPC
	t.Run("TestInitializeConnections", func(t *testing.T) {
		t.Run("WithSuccess", func(t *testing.T) {
			var err error = services.InitializeConnections()
			assert.Nil(t, err)
		})
	})
}
