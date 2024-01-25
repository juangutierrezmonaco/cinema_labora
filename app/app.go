package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/routes"
	"github.com/labora/labora-golang/cinema_labora/util"
)

func Init() {
	// Load env data
	err := util.LoadEnv()
	if err != nil {
		return
	}
	envData := util.EnvData

	// Create server
	router := mux.NewRouter()
	serv, err := config.NewServer(envData.ServerData.Host, envData.ServerData.Port, router)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	go serv.Start()

	// Init DB
	config.InitDb()

	// Routes
	routes.BuildRoutes(router)

	// Wait for an in interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	serv.Close()
}
