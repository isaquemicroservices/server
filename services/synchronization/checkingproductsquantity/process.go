package checkingproductsquantity

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/isaqueveras/servers-microservices-backend/services/synchronization/checkingproductsquantity/scripts"
)

// Watcher manage the ticker's to run scripts
func Watcher() {
	var (
		errorChannel       chan error     = make(chan error, 15)
		terminateExecution chan os.Signal = make(chan os.Signal, 1)
		continueLoop       bool           = true
	)

	signal.Notify(terminateExecution, syscall.SIGINT, syscall.SIGTERM)
	go routineRecover(errorChannel, scripts.ExecuteCheckingProductsQuantity)

	for continueLoop {
		select {
		case <-terminateExecution:
			log.Println("checkingproductsquantity/process: Signal to terminate execution was received")
			continueLoop = false
		case cannelError := <-errorChannel:
			log.Println("checkingproductsquantity/process: " + cannelError.Error())
		}
	}
}

func routineRecover(errorChannel chan error, routine func(chan error)) {
	defer func() {
		if rec := recover(); rec != nil {
			const messagePanic string = "[SCRIPT] >> [PANIC] >> A panic was identified during the execution of the scripts"
			if value, ok := rec.(error); ok {
				// TODO: add messagePanic on value of error
				log.Println(value)
				errorChannel <- value
			} else {
				fmt.Printf(messagePanic+": %v\n", rec)
			}
		}
	}()

	routine(errorChannel)
}
