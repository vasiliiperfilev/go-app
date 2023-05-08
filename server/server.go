package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if player == "test" {
		fmt.Fprint(w, "20")
	}
	if player == "anothertest" {
		fmt.Fprint(w, "30")
	}
}
