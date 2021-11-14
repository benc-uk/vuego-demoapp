package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benc-uk/vuego-demoapp/server/pkg/api"
	"github.com/gorilla/mux"
)

// API is the main backend application API
type API struct {
	WeatherAPIKey string
	// Use composition and embedding to extend the API base
	api.Base
}

// HTTPError holds API JSON error
type HTTPError struct {
	Error string `json:"error"`
}

//
// Adds backend app routes
//
func (a *API) AddRoutes(router *mux.Router) {
	router.HandleFunc("/api/info", a.getInfo).Methods("GET")
	router.HandleFunc("/api/monitor", a.getMonitorMetrics).Methods("GET")
	router.HandleFunc("/api/config", a.getConfig).Methods("GET")
	router.HandleFunc("/api/weather/{lat}/{long}", a.getWeather).Methods("GET")
	a.Healthy = true
}

//
// Helper function for returning API errors
//
func apiError(resp http.ResponseWriter, code int, message string) {
	resp.WriteHeader(code)

	//message = strings.ReplaceAll(message, "\"", "'")
	errorData := &HTTPError{
		Error: message,
	}

	errorResp, err := json.Marshal(errorData)
	if err != nil {
		fmt.Printf("### ERROR! httpError unable to marshal to JSON. Message was %s\n", message)
		return
	}
	_, _ = resp.Write(errorResp)
}
