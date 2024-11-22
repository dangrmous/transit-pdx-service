package main

import (
	"log"
	"os"
	"transit-pdx-service/service"
)

func main() {
	appLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	appService := service.NewService()
	// Let's get out of main asap
	err := appService.Start(appLogger)
	if err != nil {
		appLogger.Fatal(err)
	}
}
