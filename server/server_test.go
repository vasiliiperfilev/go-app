package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/vasiliiperfilev/go-app/server"
)

func TestServer(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/players/test", nil)
	response := httptest.NewRecorder()
	server.PlayerServer(response, request)

	got := response.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
