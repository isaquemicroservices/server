package services

import (
	config "github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
)

// InitializeConnections initialize persistent connections with services backends
func InitializeConnections(config *config.Configuration) (err error) {
	if err = grpc.InitializeProductConnections(config); err != nil {
		return err
	}

	if err = grpc.InitializeAuthConnections(config); err != nil {
		return err
	}

	if err = grpc.InitializeEmailConnections(config); err != nil {
		return err
	}

	return
}
