package grpc

import (
	"context"

	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/auth/grpc/auth"
	gogrpc "google.golang.org/grpc"
)

// Auth is a base struct
type Auth struct {
	client auth.AuthClient
	ctx    context.Context
}

// NewAuthDriver creates a new driver for querying auth data using the backend provided by the connection
func NewAuthDriver(ctx context.Context, conn gogrpc.ClientConnInterface) *Auth {
	return &Auth{
		client: auth.NewAuthClient(conn),
		ctx:    ctx,
	}
}
