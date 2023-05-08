package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/vasiliiperfilev/go-app/server"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestServer(t *testing.T) {
	playerStoreStub := StubPlayerStore{
		scores: map[string]int{
			"test":        20,
			"anothertest": 30,
		},
	}
	playerServer := &server.PlayerServer{Store: &playerStoreStub}
	t.Run("GET test player score", func(t *testing.T) {
		request := newGetScoreRequest("test")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertBody(t, got, want)
	})

	t.Run("GET anothertest player score", func(t *testing.T) {
		request := newGetScoreRequest("anothertest")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)

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
