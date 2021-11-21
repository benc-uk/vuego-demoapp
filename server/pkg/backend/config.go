package backend

import (
	"encoding/json"
	"net/http"
	"os"
)

// Used by config endpoint
type configData struct {
	AuthClientID   string `json:"authClientId"`
	WeatherEnabled bool   `json:"weatherEnabled"`
}

// Route for fetching config from the server
func (a *API) getConfig(resp http.ResponseWriter, req *http.Request) {
	config := &configData{}

	config.AuthClientID = os.Getenv("AUTH_CLIENT_ID")
	_, weatherEnabled := os.LookupEnv("WEATHER_API_KEY")
	config.WeatherEnabled = weatherEnabled

	jsonResp, err := json.Marshal(config)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	// Fire JSON result back down the internet tubes
	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write(jsonResp)
}
