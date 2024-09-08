package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPhotosEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/photos", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPhotosHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", rr.Header().Get("Content-Type"), "application/json")
	}

	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned invalid JSON: %s", rr.Body.String())
	}
}

func TestPostsEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPostsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", rr.Header().Get("Content-Type"), "application/json")
	}

	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned invalid JSON: %s", rr.Body.String())
	}
}

func TestUsersEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", rr.Header().Get("Content-Type"), "application/json")
	}

	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("handler returned invalid JSON: %s", rr.Body.String())
	}
}
