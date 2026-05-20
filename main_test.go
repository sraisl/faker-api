package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetingEndpoint(t *testing.T) {
	router := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/greeting/Alice", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if body["message"] != "Hello, Alice!" {
		t.Fatalf("expected message %q, got %q", "Hello, Alice!", body["message"])
	}
}

func TestRandomEmailEndpoint(t *testing.T) {
	router := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/random-email", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	email, ok := body["email"]
	if !ok || email == "" {
		t.Fatal("expected non-empty email field in response")
	}
}

func TestRandomNameEndpoint(t *testing.T) {
	router := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/random-name", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if body["name"] == "" {
		t.Fatal("expected non-empty name")
	}
}

func TestHealthEndpoint(t *testing.T) {
	router := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if body["status"] != "OK" {
		t.Fatalf("expected status %q, got %q", "OK", body["status"])
	}
}
