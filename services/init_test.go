package services_test

import (
	"testing"
	"time"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/stretchr/testify/assert"
)

func TestServices(t *testing.T) {
	// loading config of system
	configuration.Load()

	// TestInitializeConnections function to test initializing connections with services gRPC
	t.Run("TestInitializeConnections", func(t *testing.T) {
		// WithSuccess test for initialize connection with microservice product with success
		t.Run("WithSuccess", func(t *testing.T) {
			var err error = services.InitializeConnections(configuration.Get())
			assert.Nil(t, err)
		})

		// WithError test for initialize connection with microservice product with error
		t.Run("WithError", func(t *testing.T) {
			configuration.Get().ContextWithTimeout = time.Second * 0
			var err error = services.InitializeConnections(configuration.Get())
			assert.Error(t, err)
		})
	})
}
