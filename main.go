package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/interfaces/product"
	"github.com/isaqueveras/servers-microservices-backend/middleware"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Initialize the gin
	routes := gin.New()
	// gin.SetMode(gin.ReleaseMode)

	// middlewares
	routes.Use(middleware.CORS())

	// Initialize connections with services
	if err := services.InitializeConnections(); err != nil {
		log.Fatal(err)
		return
	}

	// Group for the  version first system
	v1 := routes.Group("v1")

	// Groups for products
	product.Router(v1.Group("products"))

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
