package main

import "sync"

type InMemoryPlayerStore struct {
	mu   sync.Mutex
	wins map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.wins[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.wins[name] += 1
}

func (i *InMemoryPlayerStore) GetLeague() League {
	var league League
	for name, wins := range i.wins {
		league = append(league, Player{name, wins})
	}
	return league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{wins: map[string]int{}}
}
