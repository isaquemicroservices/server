package grpc

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/crm/auth"
	auth "github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/auth/grpc/auth"
	"github.com/isaqueveras/servers-microservices-backend/oops"
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

// CreateUser create a new user
func (a *Auth) CreateUser(in *domain.User) (err error) {
	if _, err = a.client.CreateUser(a.ctx, &auth.User{
		Name:  in.Name,
		Email: in.Email,
		Passw: in.Passw,
	}); err != nil {
		return oops.Err(err)
	}

	return
}

// Login authentication a user
func (a *Auth) Login(user *domain.Credentials) (*domain.User, error) {
	var (
		data *auth.User
		err  error
	)

	if data, err = a.client.Login(a.ctx, &auth.Credentials{
		Email: user.Email,
		Passw: user.Passw,
	}); err != nil {
		return nil, oops.Err(err)
	}

	return &domain.User{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
		Passw: data.Passw,
		Token: data.Token,
	}, nil
}
