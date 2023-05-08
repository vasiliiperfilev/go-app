package main

import (
	"log"
	"net/http"

	"github.com/vasiliiperfilev/go-app/server"
)

func main() {
	playerServer := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", playerServer))
}
