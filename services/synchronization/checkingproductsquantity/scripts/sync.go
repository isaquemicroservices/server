package scripts

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	productApp "github.com/isaqueveras/servers-microservices-backend/application/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/utils"
)

const intervalToRunTheScript time.Duration = time.Second * 30

// ExecuteCheckingProductsQuantity checking products quantity
func ExecuteCheckingProductsQuantity(errorChannel chan error) {
	var ticker = time.NewTicker(intervalToRunTheScript)

	// nolint:gosimple
	for {
		select {
		case <-ticker.C:
			checkingProductsQuantity(errorChannel)
		}
	}
}

func checkingProductsQuantity(errorChannel chan error) {
	var (
		prefixScript        string = "[SCRIPT] >> [CHECKING_PRODUCTS_QUANTITY] >> "
		maxErrorConsecutive int    = 4
		maxRequest          int    = 3

		group             = &utils.Mutex{MaxInteraction: int64(maxRequest)}
		errorConsecutive  = &utils.Mutex{}
		requestsMade      = &utils.Mutex{}
		processedProducts = &utils.Mutex{}

		ctx = context.Background()
	)

	var timeNow = time.Now()

	// TODO: implements zap log
	log.Println(prefixScript + "Initializing the script to checking products quantity")
	defer log.Println(prefixScript + "Initializing the script to checking products quantity finished in " + time.Since(timeNow).String())

	pdct, err := productApp.GetProducts(ctx)
	if err != nil {
		log.Println(prefixScript + err.Error())
		errorConsecutive.Add(1)
		errorChannel <- err
	}

	for i := range pdct.Data {
		if int(errorConsecutive.Status()) >= maxErrorConsecutive {
			errorChannel <- errors.New("Execution of checking products quantity suspended due to excessive consecutive errors")
			break
		}

		group.Wait()
		group.Add(1)
		requestsMade.Add(1)
		defer group.Dec()

		_, err := productApp.GetDetailsProduct(ctx, pdct.Data[i].ID)
		if err != nil {
			log.Println(prefixScript + err.Error())
			errorChannel <- err
			errorConsecutive.Add(1)
			return
		}

		processedProducts.Add(1)
		errorConsecutive.Done()
		group.Done()
	}

	group.WaitDone()

	// TODO: implementar zap log
	log.Println(prefixScript + "Quantity of products: " + strconv.FormatInt(int64(len(pdct.Data)), 10))
	log.Println(prefixScript + "Confirmed requests: " + strconv.FormatInt(requestsMade.Status(), 10))
	log.Println(prefixScript + "Processed products: " + strconv.FormatInt(processedProducts.Status(), 10))
	log.Println(prefixScript + "Unprocessed products: " + strconv.FormatInt(int64(len(pdct.Data))-processedProducts.Status(), 10))
}
