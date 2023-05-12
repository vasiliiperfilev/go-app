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
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
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

func TestStoreWin(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}
	playerServer := &server.PlayerServer{Store: &store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Peper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)

		gotStatus := response.Code
		wantStatus := http.StatusAccepted

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}

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
