package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer creates a new server for players
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
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
