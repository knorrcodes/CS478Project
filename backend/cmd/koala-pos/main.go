// This source file is part of the Packet Guardian project.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/lfkeitel/verbose/v5"

	"koala.pos/src/common"
	"koala.pos/src/db"
	"koala.pos/src/models/stores"
	"koala.pos/src/server"
)

var (
	configFile string
	dev        bool
	verFlag    bool
	testConfig bool

	version   = ""
	buildTime = ""
	builder   = ""
	goversion = ""
)

func init() {
	flag.StringVar(&configFile, "c", "", "Configuration file path")
	flag.BoolVar(&dev, "d", false, "Run in development mode")
	flag.BoolVar(&testConfig, "t", false, "Test main configuration file")
	flag.BoolVar(&verFlag, "version", false, "Display version information")
	flag.BoolVar(&verFlag, "v", verFlag, "Display version information")
}

func main() {
	common.SystemVersion = version

	// Parse CLI flags
	flag.Parse()

	if verFlag {
		displayVersionInfo()
		return
	}

	if configFile == "" || !common.FileExists(configFile) {
		configFile = common.FindConfigFile()
	}
	if configFile == "" {
		fmt.Println("No configuration file found")
		os.Exit(1)
	}

	if testConfig {
		testMainConfig()
		return
	}

	e := setupEnvironment()
	startShutdownWatcher(e)

	if err := common.RunSystemInits(e); err != nil {
		log.WithField("error", err).Fatal("System initialization failed")
	}

	appStores := &stores.StoreCollection{
		Product:  stores.NewProductStore(e),
		Category: stores.NewCategoryStore(e),
	}

	// Start web server
	server.NewServer(e, server.LoadRoutes(e, appStores)).Run()
}

func setupEnvironment() *common.Environment {
	var err error
	e := common.NewEnvironment(common.EnvProd)
	if dev {
		e.Env = common.EnvDev
	}

	e.Config, err = common.NewConfig(configFile)
	if err != nil {
		fmt.Printf("Error loading configuration: %s\n", err)
		os.Exit(1)
	}

	log.WithField("config_file", configFile).Debug("Configuration loaded")

	e.DB, err = db.NewDatabaseAccessor(e)
	if err != nil {
		log.WithField("error", err).Fatal("Error loading database")
	}
	log.WithFields(log.Fields{
		"address": e.Config.Database.Address,
	}).Debug("Loaded database")

	return e
}

func startShutdownWatcher(e *common.Environment) {
	c := e.SubscribeShutdown()
	go func(e *common.Environment) {
		<-c
		if err := e.DB.Close(); err != nil {
			log.WithField("error", err).Warning("Error closing database")
		}
		log.Notice("Shutting down...")
		time.Sleep(2)
	}(e)
}

func displayVersionInfo() {
	fmt.Printf(`Packet Guardian - (C) 2019 The Koala POS Authors

Component:   API Server
Version:     %s
Built:       %s
Compiled by: %s
Go version:  %s
`, version, buildTime, builder, goversion)
}

func testMainConfig() {
	_, err := common.NewConfig(configFile)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Configuration looks good")
}
