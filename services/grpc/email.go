package grpc

import (
	"context"
	"sync"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	emailConnection *gogrpc.ClientConn
	emailMutex      sync.Mutex
)

// InitializeEmailConnections open connections with service gRPC
func InitializeEmailConnections(config *configuration.Configuration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configuration.Get().ContextWithTimeout)
	defer cancel()

	if emailConnection, err = gogrpc.DialContext(ctx, config.MSEmail.Address, gogrpc.WithInsecure(), gogrpc.WithBlock()); err != nil {
		return err
	}

	return
}

// GetEmailConnection return an active connection with the email backend
func GetEmailConnection() *gogrpc.ClientConn {
	emailMutex.Lock()
	defer emailMutex.Unlock()

	if state := emailConnection.GetState(); state != connectivity.Ready {
		_ = InitializeEmailConnections(configuration.Get())
	}

	return emailConnection
}
