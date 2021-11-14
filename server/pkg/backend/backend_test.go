package backend

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/benc-uk/vuego-demoapp/server/pkg/api"
)

var backendAPI *API

func init() {
	backendAPI = &API{
		WeatherAPIKey: "",
		Base: api.Base{
			Healthy: true,
			Version: "3.0.0",
			Name:    "VueGo-DemoApp-API",
		},
	}
}

//
// Test TestApiInfoRoute
//
func TestApiInfoRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/info", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(backendAPI.getInfo)
	handler.ServeHTTP(rr, req)

	// Check resp code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "hostname"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("TestApiInfoRoute returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//
// Test TestWeatherRoute
//
func TestWeatherRoute(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/weather", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(backendAPI.getWeather)
	handler.ServeHTTP(rr, req)

	// Check resp code - it will fail!
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("TestWeatherRoute returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}

//
// Test TestConfigRoute
//
func TestConfigRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/config", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(backendAPI.getConfig)
	handler.ServeHTTP(rr, req)

	// Check resp code - it will fail!
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("TestWeatherRoute returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "authClientId"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("TestApiInfoRoute returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
