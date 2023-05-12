package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vasiliiperfilev/go-app/server"
	"github.com/vasiliiperfilev/go-app/store"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	server := server.PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertValue(t, response.Code, http.StatusOK)
	assertValue(t, response.Body.String(), "3")
}
