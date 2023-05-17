package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemStore struct {
	database io.Writer
	league   League
}

func (f *FileSystemStore) GetLeague() League {
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		newPlayer := Player{Name: name, Wins: 1}
		f.league = append(f.league, newPlayer)
	}

	json.NewEncoder(f.database).Encode(f.league)
}

func NewFileSystemStore(db *os.File) *FileSystemStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemStore{
		database: &tape{db},
		league:   league,
	}
}
