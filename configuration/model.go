package configuration

import "time"

// Configuration main configuration struct
type Configuration struct {
	NameServer         string        `json:"name_application"`
	ServerAddress      string        `json:"server_address"`
	ContextWithTimeout time.Duration `json:"context_timeout"`
	MSProduct          microservice  `json:"microservice_product"`
	MSAuthentication   microservice  `json:"microservice_authentication"`
}

type microservice struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
