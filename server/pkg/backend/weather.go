package backend

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Route to proxy weather data from api.openweathermap.org
func (a *API) getWeather(resp http.ResponseWriter, req *http.Request) {
	if a.WeatherAPIKey == "" {
		apiError(resp, http.StatusNotImplemented, "Feature disabled, WEATHER_API_KEY is not set")
		return
	}

	vars := mux.Vars(req)

	// Fetch fetch weather data from OpenWeather API
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", vars["lat"], vars["long"], a.WeatherAPIKey)
	apiResp, err := http.Get(url)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	if apiResp.StatusCode != 200 {
		apiError(resp, apiResp.StatusCode, "OpenWeather API error: "+apiResp.Status)
		return
	}
	// We simply proxy the result back to the client
	body, _ := ioutil.ReadAll(apiResp.Body)

	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write(body)
}
