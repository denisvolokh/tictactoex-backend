package main

import (
	"encoding/json"
	"go-tic-tac-toe-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestGetGamesEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/games", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getGamesHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	var games []models.Game
	json.NewDecoder(rr.Body).Decode(&games)

	expected := 6
	if len(games) != expected {
		t.Errorf("handler returned unexpected body: got %v, expected %v", len(games), expected)
	}
}

func TestPostGamesEndpoint(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/games", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postGamesHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("POST handler returned wrong status code: got %v, expected %v", status, http.StatusCreated)
	}

	var game models.Game
	json.NewDecoder(rr.Body).Decode(&game)

	_, err = uuid.Parse(game.ID)
	if err != nil {
		t.Errorf("POST handler returned unexpected body: got %v, expected %v", game, err)
	}
}

func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}
