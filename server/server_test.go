package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/vasiliiperfilev/go-app/server"
)

func TestServer(t *testing.T) {
	t.Run("GET test player score", func(t *testing.T) {
		request := newGetScoreRequest("test")
		response := httptest.NewRecorder()
		server.PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		assertBody(t, got, want)
	})

	t.Run("GET anothertest player score", func(t *testing.T) {
		request := newGetScoreRequest("anothertest")
		response := httptest.NewRecorder()
		server.PlayerServer(response, request)

		got := response.Body.String()
		want := "30"

		assertBody(t, got, want)
	})

}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
