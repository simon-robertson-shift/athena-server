package main

import (
	"athena/server"
	"athena/server/database"
	"athena/server/database/versions"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	development := false

	for _, value := range os.Args {
		if value == "dev" {
			development = true
		}
	}

	if development {
		log.Println("[running in development mode]")
	} else {
		log.Println("[running in production mode]")
	}

	log.Println("Loading the environment ...")

	if development {
		godotenv.Load(".env.development")
	} else {
		godotenv.Load(".env.production")
	}

	log.Println("Connecting to the database ...")

	if err := database.Connect(development); err != nil {
		log.Println("Failed to connect")
		log.Println(err)
		return
	}

	defer database.Disconnect()

	log.Println("Checking database versions ...")
	versions.Sync()

	log.Println("Loading constant store data ...")
	// constants.Load()

	log.Println("Starting the server ...")
	server.Start()
}
