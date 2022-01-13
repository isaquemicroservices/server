package auth

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/crm/auth"
	infra "github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/auth"
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
)

// Create contains the logic to create a user
func Create(ctx context.Context, in *User) (err error) {
	var (
		conn = grpc.GetAuthConnection()
		repo = infra.New(ctx, conn)
	)

	defer conn.Close()

	var user = &domain.User{
		Name:  in.Name,
		Email: in.Email,
		Passw: in.Passw,
	}

	if err = repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

// Login contains the logic to authentication a user
func Login(ctx context.Context, in *CredentialsReq) (res *User, err error) {
	var (
		conn = grpc.GetAuthConnection()
		repo = infra.New(ctx, conn)
		data *domain.User
	)

	defer conn.Close()

	var user = &domain.Credentials{
		Email: in.Email,
		Passw: in.Passw,
	}

	if data, err = repo.Login(user); err != nil {
		return nil, err
	}

	return &User{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
		Passw: data.Passw,
		Token: data.Token,
	}, nil
}
