package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store := NewFileSystemStore(database)
	playerServer := NewPlayerServer(store)
	player := "Pepper"

	playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newGetScoreRequest(player))
		assertValue(t, response.Code, http.StatusOK)

		assertValue(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newLeagueRequest())
		assertValue(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{Name: "Pepper", Wins: 3},
		}
		assertValue(t, got, want)
	})
}
