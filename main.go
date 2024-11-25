package main

import (
	"github.com/dangrmous/transit-pdx-service/service"
	"log"
	"os"
)

func main() {
	appLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	appService := service.New()
	// Let's get out of main asap
	err := appService.Start(appLogger)
	if err != nil {
		appLogger.Fatal(err)
	}
}
