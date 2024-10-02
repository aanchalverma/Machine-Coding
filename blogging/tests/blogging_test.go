package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aanchalverma/Machine-Coding/blogging/routes"
)

func TestCreatePost(t *testing.T) {
	// Create a new HTTP POST request
	var jsonStr = []byte(`{"title":"Test Post", "content":"This is a test post", "author":"John Doe"}`)
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer <your JWT token>")

	// Use httptest to create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HandleCreatePost)
	handler.ServeHTTP(rr, req)

	// Check if the status code is 201 (Created)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestGetPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HandleGetPosts)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `[{"title":"Test Post","content":"This is a test post","author":"John Doe"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
