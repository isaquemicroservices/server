package product

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	router := gin.Default()

	if err := services.InitializeConnections(); err != nil {
		log.Fatal(err)
	}

	// Initializing routes to crm/product
	Router(router.Group("v1/crm/products"))
	RouterWithID(router.Group("/v1/crm/product"))

	// Test of route to get all products
	t.Run("TestGetAllProducts", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/v1/crm/products", nil)
		router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, 200, w.Code)
	})

	// Test of route to get details of product
	t.Run("TestGetDetailsProduct", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/v1/crm/product/16", nil)
		router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, 200, w.Code)
	})
}
