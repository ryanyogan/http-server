package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore is a temporary data-base for us to work with
type InMemoryPlayerStore struct{}

// GetPlayerScore takes a name as a string and returns the score as an
// integer.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin takes a players name and adds it to the store.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	// I DO NOTHING
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
