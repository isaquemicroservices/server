package services

import (
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
)

// InitializeConnections initialize persistent connections with services backends
func InitializeConnections() (err error) {
	if err = grpc.InitializeProductConnections(); err != nil {
		return err
	}

	return
}
