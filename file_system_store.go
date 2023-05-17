package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemStore struct {
	database *json.Encoder
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

	f.database.Encode(f.league)
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	file.Seek(0, 0)
	info, err := file.Stat()

	if err != nil {
		return nil, fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, err
	}
	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}
