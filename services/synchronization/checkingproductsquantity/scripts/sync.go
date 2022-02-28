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

const intervalToRunTheScript time.Duration = time.Second * 10

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
	)

	var timeNow = time.Now()

	// TODO: implements zap log
	log.Println(prefixScript + "Initializing the script to checking products quantity")
	defer log.Println(prefixScript + "Initializing the script to checking products quantity finished in " + time.Since(timeNow).String())

	var ctx = context.Background()
	listProducts, err := productApp.ListAllProductsWithMinimumQuantity(ctx)
	if err != nil {
		log.Println(prefixScript + err.Error())
		errorConsecutive.Add(1)
		errorChannel <- err
	}

	if int(errorConsecutive.Status()) >= maxErrorConsecutive {
		errorChannel <- errors.New("Execution of checking products quantity suspended due to excessive consecutive errors")
	}

	group.Wait()
	group.Add(1)
	requestsMade.Add(1)

	// TODO: send email for admin's the of list of products with minimum quantity
	go sendEmail(ctx, listProducts.Data, group, processedProducts, errorConsecutive)

	group.WaitDone()

	// TODO: implementar zap log
	log.Println(prefixScript + "Quantity of products: " + strconv.FormatInt(int64(len(listProducts.Data)), 10))
	log.Println(prefixScript + "Confirmed requests: " + strconv.FormatInt(requestsMade.Status(), 10))
	log.Println(prefixScript + "Processed products: " + strconv.FormatInt(processedProducts.Status(), 10))
	log.Println(prefixScript + "Unprocessed products: " + strconv.FormatInt(int64(len(listProducts.Data))-processedProducts.Status(), 10))
}

func sendEmail(_ context.Context, products []productApp.Product, group, processedProducts, errorConsecutive *utils.Mutex) {
	defer group.Dec()

	var listProductsEmail []string

	for i := range products {
		listProductsEmail = append(listProductsEmail, "#"+strconv.FormatInt(*products[i].ID, 10)+" "+*products[i].Name+" ("+strconv.FormatInt(*products[i].Amount, 10)+" un.)")
		processedProducts.Add(1)
	}

	for _, v := range listProductsEmail {
		log.Println(v)
	}

	errorConsecutive.Done()
}
