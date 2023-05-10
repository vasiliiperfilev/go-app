package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestGetScore(t *testing.T) {
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

		gotBody := response.Body.String()
		wantBody := "20"

		gotStatus := response.Code
		wantStatus := http.StatusOK

		assertValue(t, gotBody, wantBody)
		assertValue(t, gotStatus, wantStatus)
	})

	t.Run("GET anothertest player score", func(t *testing.T) {
		request := newGetScoreRequest("anothertest")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)

		gotBody := response.Body.String()
		wantBody := "30"

		gotStatus := response.Code
		wantStatus := http.StatusOK

		assertValue(t, gotBody, wantBody)
		assertValue(t, gotStatus, wantStatus)
	})

	t.Run("GET missing player score", func(t *testing.T) {
		request := newGetScoreRequest("missing")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)

		gotStatus := response.Code
		wantStatus := http.StatusNotFound

		assertValue(t, gotStatus, wantStatus)
	})
}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertValue[T any](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
