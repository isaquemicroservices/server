package configuration

import "time"

const (
	// port of server
	PortServer = "localhost:8080"
	// url of product service
	ProductURL = "localhost:50051"
	// Timeout to context
	ContextWithTimeout = time.Second * 2
)
