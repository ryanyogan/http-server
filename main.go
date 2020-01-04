package main

import (
	"log"
	"net/http"
)

// -- Temp memory store until we need a DB -- //

// NewInMemoryPlayerStore returns a new instance of the Memory backed
// data-store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore is a temporary data-base for us to work with
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore takes a name as a string and returns the score as an
// integer.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin takes a players name and adds it to the store.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// -- End of Temp memory store block //

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
