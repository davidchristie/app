package main

import (
	"log"
	"os"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/database"
	"github.com/davidchristie/app/services/app/http"
)

var server http.Server

var exit = os.Exit

func init() {
	initServer()
}

func initServer() {
	config, err := config.LoadConfig()
	if err != nil {
		logFatal(err)
		return
	}
	_, err = database.NewConnection(config)
	if err != nil {
		log.Println("Error connecting to database:", err) // TODO
	}
	auth := auth.NewAuth()
	server = http.NewServer(config, auth)
}

func logFatal(v ...interface{}) {
	log.Print(v...)
	exit(1)
}

func main() {
	logFatal(server.Start())
}
