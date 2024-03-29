package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemStore(db)
	if err != nil {
		log.Fatal("Can't establish db")
	}
	server := &PlayerServer{Store: store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
