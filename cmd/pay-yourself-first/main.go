package main

import (
	"log"
	"net/http"

	"github.com/catzkorn/pay-yourself-first/internal/database"
	"github.com/catzkorn/pay-yourself-first/internal/server"
)

func main() {

	database, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	server := server.NewServer(database)
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
