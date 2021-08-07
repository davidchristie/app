package main

import (
	"log"
	"os"

	"github.com/davidchristie/app/services/app/auth"
	"github.com/davidchristie/app/services/app/config"
	"github.com/davidchristie/app/services/app/database"
	"github.com/davidchristie/app/services/app/http"
	"github.com/davidchristie/app/services/app/repositories"
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
	db, err := database.NewConnection(config)
	if err != nil {
		log.Println("Error connecting to database:", err) // TODO
	}
	userRepository := repositories.NewUserRepository(db)
	accountRepository := repositories.NewAccountRepository(db)
	sessionRepository := repositories.NewSessionRepository(db)
	auth := auth.NewAuth(
		config,
		userRepository,
		accountRepository,
		sessionRepository,
	)
	server = http.NewServer(config, auth)
}

func logFatal(v ...interface{}) {
	log.Print(v...)
	exit(1)
}

func main() {
	logFatal(server.Start())
}
