package store

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
