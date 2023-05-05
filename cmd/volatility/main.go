package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/ATMackay/go-quantifier/common"
	"github.com/ATMackay/go-quantifier/service/volatility"
)

const envPrefix = "VOL"

var (
	configFile    string
	configFilePtr = flag.String("config", "vol.yml", "path to config file")
)

func init() {
	// Parse flag containing path to config file
	flag.Parse()
	if configFilePtr != nil {
		configFile = *configFilePtr
	}
}

func main() {

	// Step 1) Initialise config variables
	//        * flags
	//
	// Step 2) Build base application
	//        * HTTPS server
	//        * Database
	//        * Remote data fetcher
	// Step 3) Start application

	var config volatility.Config
	if err := common.ParseYAMLConfig(configFile, &config, envPrefix); err != nil {
		log.Fatal(err)
	}

	volatilityService, err := volatility.BuildService(config)
	if err != nil {
		log.Fatal(err)
	}

	volatilityService.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	volatilityService.Stop()
}
