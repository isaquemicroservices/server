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
	SecretKey          string        `json:"jwt_secret_key"`
}

type microservice struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

// Session session model
type Session struct {
	Administrator *bool   `json:"administrator,omitempty"`
	Name          *string `json:"name,omitempty"`
	jwt.StandardClaims
}
