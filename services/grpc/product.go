package grpc

import (
	"log"
	"sync"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	productConnection *grpc.ClientConn
	productMutex      sync.Mutex
	err               error
)

// InitializeProductConnections open connections with service gRPC
func InitializeProductConnections() (err error) {
	if productConnection, err = grpc.Dial(configuration.ProductURL, grpc.WithInsecure()); err != nil {
		return err
	}

	return
}

// GetProductConnection return an active connection with the product backend
func GetProductConnection() *grpc.ClientConn {
	productMutex.Lock()
	defer productMutex.Unlock()

	if state := productConnection.GetState(); state != connectivity.Ready {
		if err = InitializeProductConnections(); err != nil {
			log.Fatal(err)
			return nil
		}
	}

	return productConnection
}
