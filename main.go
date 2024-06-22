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
		log.Println("DEVELOPMENT")
	} else {
		log.Println("PRODUCTION")
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

	log.Println("Running")
	server.Start()
}
