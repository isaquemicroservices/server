package configuration

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Configuration main configuration struct
type Configuration struct {
	NameServer         string        `json:"name_application"`
	ServerAddress      string        `json:"server_address"`
	ContextWithTimeout time.Duration `json:"context_timeout"`
	MSProduct          microservice  `json:"microservice_product"`
	MSAuthentication   microservice  `json:"microservice_authentication"`
	MSEmail            microservice  `json:"microservice_email"`
	SecretKey          string        `json:"jwt_secret_key"`
}

type microservice struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

// Session session model
type Session struct {
	Name       *string    `json:"name,omitempty"`
	Email      *string    `json:"email,omitempty"`
	Permission *UserLevel `json:"permission,omitempty"`
	jwt.StandardClaims
}

// UserLevel permission model of user
type UserLevel struct {
	IsAdmin *bool   `json:"is_admin,omitempty"`
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
}
