package product

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/application/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/isaqueveras/servers-microservices-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	router := gin.Default()

	// loading config of system
	configuration.Load()

	if err := services.InitializeConnections(configuration.Get()); err != nil {
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

	// Test of route to create product
	t.Run("TestCreateProduct", func(t *testing.T) {
		w := httptest.NewRecorder()

		var productJson, err = json.Marshal(product.Product{
			Name:        utils.GetPointerString("Milk"),
			Description: utils.GetPointerString("Milk is essentially an emulsion of fat and protein in water, along with dissolved sugar (carbohydrate), minerals, and vitamins."),
			Price:       utils.GetPointerFloat64(4.59),
		})

		assert.Nil(t, err)

		req, err := http.NewRequest("POST", "/v1/crm/products", bytes.NewBuffer(productJson))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Nil(t, err)
		assert.Equal(t, 201, w.Code)
	})
}
