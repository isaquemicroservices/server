package grpc

import (
	"sync"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	productConnection *gogrpc.ClientConn
	productMutex      sync.Mutex
)

// InitializeProductConnections open connections with service gRPC
func InitializeProductConnections() (err error) {
	if productConnection, err = gogrpc.Dial(configuration.ProductURL, gogrpc.WithInsecure()); err != nil {
		return err
	}

	return
}

// GetProductConnection return an active connection with the product backend
func GetProductConnection() gogrpc.ClientConnInterface {
	productMutex.Lock()
	defer productMutex.Unlock()

	if state := productConnection.GetState(); state != connectivity.Ready {
		_ = InitializeProductConnections()
	}

	return productConnection
}
