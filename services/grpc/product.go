package grpc

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), configuration.Get().ContextWithTimeout)
	defer cancel()

	if productConnection, err = gogrpc.DialContext(ctx, config.ProductAddress, gogrpc.WithInsecure(), gogrpc.WithBlock()); err != nil {
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
