package main

import (
	"log"

	"github.com/HemlockPham7/server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	log.Fatal("connected to database")
}
