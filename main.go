package main

import (
	"log"
	"net/http"

	"github.com/vasiliiperfilev/go-app/server"
	"github.com/vasiliiperfilev/go-app/store"
)

func main() {
	server := &server.PlayerServer{Store: store.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
