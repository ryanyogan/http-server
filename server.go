package main

import "net/http"

import "fmt"

// PlayerServer creates a new server for players
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
