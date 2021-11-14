package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/benc-uk/vuego-demoapp/server/pkg/api"
	"github.com/benc-uk/vuego-demoapp/server/pkg/backend"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	fmt.Println("### 🚀 Go server backend and REST API is starting...")

	// Get server PORT setting or use 4000 as default
	serverPort := "4000"
	if envValue, envSet := os.LookupEnv("PORT"); envSet {
		serverPort = envValue
	}

	// Get CONTENT_DIR setting for static content or default
	contentDir := "."
	if envValue, envSet := os.LookupEnv("CONTENT_DIR"); envSet {
		contentDir = envValue
	}

	// Enable optional weather feature
	weatherAPIKey := ""
	if envValue, envSet := os.LookupEnv("WEATHER_API_KEY"); envSet {
		fmt.Println("### 🌞 Weather API feature enabled")
		weatherAPIKey = envValue
	}

	if len(os.Getenv("AUTH_CLIENT_ID")) > 0 {
		fmt.Printf("### 🔐 Azure AD configured with client id: %s\n", os.Getenv("AUTH_CLIENT_ID"))
	}

	// Routing using mux
	router := mux.NewRouter()

	backendAPI := &backend.API{
		WeatherAPIKey: weatherAPIKey,
		Base: api.Base{
			Healthy: true,
			Version: "3.0.0",
			Name:    "VueGo-DemoApp-API",
		},
	}

	// Bind application routes to the router
	backendAPI.AddRoutes(router)

	// Add logging, CORS, health & metrics middleware
	backendAPI.AddLogging(router)
	backendAPI.AddCORS([]string{"*"}, router)
	backendAPI.AddMetrics(router, "/api")
	backendAPI.AddHealth(router, "/api")
	backendAPI.AddStatus(router, "/api")

	// Add static SPA hosting
	spa := spaHandler{staticPath: contentDir, indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	server := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           router,
		Addr:              ":" + serverPort,
	}

	// Start server
	fmt.Printf("### 🌐 HTTP server listening on %v\n", serverPort)
	fmt.Printf("### 📁 Serving static content from '%v'\n", contentDir)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
