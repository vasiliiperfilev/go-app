package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
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

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}

func NewFileSystemStore(db io.ReadWriteSeeker) *FileSystemStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemStore{
		database: db,
		league:   league,
	}
}
