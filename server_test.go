package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func TestGetScore(t *testing.T) {
	playerStoreStub := StubPlayerStore{
		scores: map[string]int{
			"test":        20,
			"anothertest": 30,
		},
	}
	playerServer := NewPlayerServer(&playerStoreStub)
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
	playerStore := StubPlayerStore{
		scores: map[string]int{},
	}
	playerServer := NewPlayerServer(&playerStore)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Peper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)

		gotStatus := response.Code
		wantStatus := http.StatusAccepted

		if len(playerStore.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(playerStore.winCalls), 1)
		}

		if playerStore.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", playerStore.winCalls[0], player)
		}

		assertValue(t, gotStatus, wantStatus)
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}

	playerStore := StubPlayerStore{nil, nil, wantedLeague}
	playerServer := NewPlayerServer(&playerStore)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response.Body)

		assertValue(t, response.Code, http.StatusOK)
		assertValue(t, got, wantedLeague)
		if response.Result().Header.Get("content-type") != JsonContentType {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
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
