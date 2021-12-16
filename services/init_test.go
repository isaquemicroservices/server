package services_test

import (
	"testing"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/stretchr/testify/assert"
)

func TestServices(t *testing.T) {
	// loading config of system
	configuration.Load()

	// TestInitializeConnections function to test initializing connections with services gRPC
	t.Run("TestInitializeConnections", func(t *testing.T) {
		t.Run("WithSuccess", func(t *testing.T) {
			var err error = services.InitializeConnections(configuration.Get())
			assert.Nil(t, err)
		})
	})
}
