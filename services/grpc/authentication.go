package grpc

import (
	"context"
	"sync"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	authConnection *gogrpc.ClientConn
	authMutex      sync.Mutex
)

// InitializeAuthConnections open connections with service gRPC
func InitializeAuthConnections(config *configuration.Configuration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configuration.Get().ContextWithTimeout)
	defer cancel()

	if authConnection, err = gogrpc.DialContext(ctx, config.MSAuthentication.Address, gogrpc.WithInsecure(), gogrpc.WithBlock()); err != nil {
		return err
	}

	return
}

// GetAuthConnection return an active connection with the authentication backend
func GetAuthConnection() *gogrpc.ClientConn {
	authMutex.Lock()
	defer authMutex.Unlock()

	if state := authConnection.GetState(); state != connectivity.Ready {
		_ = InitializeAuthConnections(configuration.Get())
	}

	return authConnection
}
