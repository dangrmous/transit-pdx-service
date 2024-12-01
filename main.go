package main

import (
	"github.com/dangrmous/transit-pdx-service/config"
	"github.com/dangrmous/transit-pdx-service/service"
	"github.com/dangrmous/transit-pdx-service/trimet"
	"log"
	"os"
)

func main() {
	appLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	appService := service.New(appLogger)
	//config := config.NewConfig()
	//trimetClient := trimet.NewClient()

	// Let's get out of main asap
	err := appService.Start(appLogger)
	if err != nil {
		appLogger.Fatal(err)
	}
}
