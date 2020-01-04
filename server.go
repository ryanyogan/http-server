package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer creates a new server for players
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Ryan" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprint(w, "20")
		return
	}
}
