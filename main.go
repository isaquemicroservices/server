package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/interfaces/crm"
	"github.com/isaqueveras/servers-microservices-backend/interfaces/crm/auth"
	"github.com/isaqueveras/servers-microservices-backend/middleware"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/isaqueveras/servers-microservices-backend/services/synchronization"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Initializing the gin
	routes := gin.New()

	// loading config of system
	configuration.Load()

	// Middlewares
	routes.Use(middleware.CORS())

	// Initialize connections with services
	if err := services.InitializeConnections(configuration.Get()); err != nil {
		log.Fatal("Was not possible to initialize connections with integrated systems", err)
		return
	}

	// Initializing synchronization with scripts
	synchronization.InitSynchronization()

	// Group of routes to the version first system
	v1 := routes.Group("v1")
	crm.Router(v1.Group("crm"))

	// Router for authorization
	auth.RouterWithAuth(v1.Group("auth"))
	auth.RouterWithoutAuth(v1.Group("auth"))

	grupoErro := errgroup.Group{}
	grupoErro.Go(func() error {
		return endless.ListenAndServe(configuration.Get().ServerAddress, routes)
	})

	// Initialize the server
	if err := grupoErro.Wait(); err != nil {
		log.Fatalf(err.Error())
		return
	}
}
