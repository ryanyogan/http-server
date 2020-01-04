package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore represents an interface for interracting with the
// data-store
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer represents the PlayerStore reference, this interface
// allows for any store that can use the io.Writer interface (most of Go)
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP will take a request, and return a response via the `ResponseWriteer`
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// GetPlayerScore takes a player name and returns their score as a string.
func GetPlayerScore(name string) string {
	if name == "Ryan" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
