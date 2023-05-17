package main

import (
	"strings"
	"testing"
)

func TestFileStore(t *testing.T) {
	t.Run("it returns league", func(t *testing.T) {
		data := strings.NewReader(`[
			{"Name": "Mark", "Wins": 1},
			{"Name": "Bob", "Wins": 2}]`)
		fsStore := NewFileSystemStore(data)

		got := fsStore.GetLeague()
		want := []Player{
			{Name: "Mark", Wins: 1},
			{Name: "Bob", Wins: 2},
		}

		assertValue(t, got, want)
		// second read
		got = fsStore.GetLeague()
		assertValue(t, got, want)
	})

	t.Run("it gets player score", func(t *testing.T) {
		data := strings.NewReader(`[
			{"Name": "Mark", "Wins": 1},
			{"Name": "Bob", "Wins": 2}]`)
		fsStore := NewFileSystemStore(data)

		got := fsStore.GetPlayerScore("Mark")
		want := 1

		assertValue(t, got, want)
	})
}
