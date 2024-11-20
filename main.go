package main

import (
	"log"
	"os"
	"transit-pdx-service/service"
)

func main() {
	appLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	appService := service.NewService()
	err := appService.Start(appLogger)
	if err != nil {
		appLogger.Fatal(err)
	}
}
