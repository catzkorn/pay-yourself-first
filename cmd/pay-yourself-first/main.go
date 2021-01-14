package main

import (
	"log"
	"net/http"

	"github.com/catzkorn/pay-yourself-first/budget"
	"github.com/catzkorn/pay-yourself-first/server"
)

func main() {

	database, err := budget.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	server := server.NewServer(database)
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
