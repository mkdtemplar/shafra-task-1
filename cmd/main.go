package main

import (
	"log"
	"shafra-task-1/api"
	"shafra-task-1/internal/database/db"
	"shafra-task-1/utils"
	"time"
)

func main() {
	start := time.Now()

	config, err := utils.LoadConfig(".")
	conn := db.ConnectToPostgres(config)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewUserRepo(conn)
	newServer, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Now().Sub(start)
	log.Printf("Server took %s", elapsed)
	err = newServer.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
