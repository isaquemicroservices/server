package synchronization

import "github.com/isaqueveras/servers-microservices-backend/services/synchronization/checkingproductsquantity"

// InitSynchronization initializing synchronization of services
func InitSynchronization() {
	checkingproductsquantity.InitSynchronizationCheckingProductsQuantity()
}
