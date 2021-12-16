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
func InitializeProductConnections(config *configuration.Configuration) (err error) {
	if productConnection, err = gogrpc.Dial(config.ProductAddress, gogrpc.WithInsecure(), gogrpc.WithBlock()); err != nil {
		return err
	}

	return
}

// GetProductConnection return an active connection with the product backend
func GetProductConnection() *gogrpc.ClientConn {
	productMutex.Lock()
	defer productMutex.Unlock()

	if state := productConnection.GetState(); state != connectivity.Ready {
		_ = InitializeProductConnections(configuration.Get())
	}

	return productConnection
}
