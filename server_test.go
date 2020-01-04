package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Ryan":  20,
			"Floyd": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Ryan's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Ryan", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, recieved %q, and wanted %q", got, want)
	}
}
