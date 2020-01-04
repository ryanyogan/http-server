package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore is a temporary data-base for us to work with
type InMemoryPlayerStore struct{}

// GetPlayerScore takes a name as a string and returns the score as an
// integer.
func (u *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
