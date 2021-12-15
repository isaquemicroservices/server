package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/interfaces/crm"
	"github.com/isaqueveras/servers-microservices-backend/middleware"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Initialize the gin
	routes := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// Middlewares
	routes.Use(middleware.CORS())

	// Initialize connections with services
	if err := services.InitializeConnections(); err != nil {
		log.Fatal("Was not possible to initialize connections with integrated systems", err)
		return
	}

	// Group of routes to the version first system
	v1 := routes.Group("v1")
	crm.Router(v1.Group("crm"))

	grupoErro := errgroup.Group{}
	grupoErro.Go(func() error {
		return endless.ListenAndServe(configuration.PortServer, routes)
	})

	// Initialize the server
	if err := grupoErro.Wait(); err != nil {
		log.Fatalf(err.Error())
		return
	}
}
