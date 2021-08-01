package main

import (
	"log"
	"os"

	"github.com/davidchristie/app/services/server/auth"
	"github.com/davidchristie/app/services/server/config"
	"github.com/davidchristie/app/services/server/http"
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
