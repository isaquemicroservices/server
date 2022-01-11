package product

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/crm/auth"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/auth/grpc"
	gogrpc "google.golang.org/grpc"
)

// repository is the base struct for repository
type repository struct {
	grpcData *grpc.Auth
}

// New initializes a repository
func New(ctx context.Context, conn gogrpc.ClientConnInterface) domain.IAuth {
	return &repository{
		grpcData: grpc.NewAuthDriver(ctx, conn),
	}
}
