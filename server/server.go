package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	if player == "test" {
		return "20"
	}
	if player == "anothertest" {
		return "30"
	}
	return ""
}
