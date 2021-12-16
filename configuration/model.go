package configuration

import "time"

// Configuration main configuration struct
type Configuration struct {
	NameServer         string        `json:"name_application"`
	ServerAddress      string        `json:"server_address"`
	ProductAddress     string        `json:"product_address"`
	ContextWithTimeout time.Duration `json:"context_timeout"`
}
