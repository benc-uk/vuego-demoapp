package main

import (
	"testing"
  "net/http/httptest"
	"net/http"
	"strings"
)

//
// Test TestApiInfoRoute
//
func TestApiInfoRoute(t *testing.T) {
	routes := Routes{
		contentDir:    "",
		disableCORS:   true,
		darkskyAPIKey: "",
		ipstackAPIKey: "",
	}

	req, err := http.NewRequest("GET", "/api/info", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.apiInfoRoute)
	handler.ServeHTTP(rr, req)

	// Check resp code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "hostname"
	if !strings.Contains (rr.Body.String(), expected) {
		t.Errorf("TestApiInfoRoute returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//
// Test TestWeatherRoute
//
func TestWeatherRoute(t *testing.T) {
	routes := Routes{
		contentDir:    "",
		disableCORS:   true,
		darkskyAPIKey: "",
		ipstackAPIKey: "",
	}

	req, err := http.NewRequest("GET", "/api/info", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.weatherRoute)
	handler.ServeHTTP(rr, req)

	// Check resp code - it will fail!
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("TestWeatherRoute returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//
// Test TestSpaRoute
//
func TestSpaRoute(t *testing.T) {
	routes := Routes{
		contentDir:    "",
		disableCORS:   true,
		darkskyAPIKey: "",
		ipstackAPIKey: "",
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.spaIndexRoute)
	handler.ServeHTTP(rr, req)

	// Check resp code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("TestSpaRoute returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}