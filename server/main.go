package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benc-uk/go-starter/pkg/envhelper"
	"github.com/gorilla/mux"
)

var contentDir string

func main() {
	// Get server PORT setting or default
	serverPort := envhelper.GetEnvString("PORT", "4000")
	// Get CONTENT_DIR setting for static content or default
	contentDir = envhelper.GetEnvString("CONTENT_DIR", ".")

	// API keys
	darkskyAPIKey := envhelper.GetEnvString("WEATHER_API_KEY", "")
	ipstackAPIKey := envhelper.GetEnvString("IPSTACK_API_KEY", "")

	fmt.Println("### 🚀 Go server backend and REST API is starting...")

	if len(darkskyAPIKey) > 0 && len(ipstackAPIKey) > 0 {
		fmt.Println("### 🌞 Weather API enabled with WEATHER_API_KEY & IPSTACK_API_KEY")
	}

	if len(os.Getenv("AUTH_CLIENT_ID")) > 0 {
		fmt.Printf("### 🔐 Azure AD configured with client id: %s\n", os.Getenv("AUTH_CLIENT_ID"))
	}

	// Routing using mux
	muxrouter := mux.NewRouter()
	routes := Routes{
		contentDir:    contentDir,
		disableCORS:   true,
		darkskyAPIKey: darkskyAPIKey,
		ipstackAPIKey: ipstackAPIKey,
	}

	// API routes
	muxrouter.HandleFunc("/api/info", routes.apiInfoRoute)
	muxrouter.HandleFunc("/api/metrics", routes.apiMetricsRoute)
	muxrouter.HandleFunc("/api/weather", routes.weatherRoute)
	muxrouter.HandleFunc("/api/config", routes.configRoute)

	// Handle static content, we have to explicitly put our top level dirs in here
	// - otherwise the NotFoundHandler will catch them
	fileServer := http.FileServer(http.Dir(contentDir))
	muxrouter.PathPrefix("/js").Handler(http.StripPrefix("/", fileServer))
	muxrouter.PathPrefix("/css").Handler(http.StripPrefix("/", fileServer))
	muxrouter.PathPrefix("/img").Handler(http.StripPrefix("/", fileServer))
	muxrouter.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/", fileServer))

	// EVERYTHING else redirect to index.html - that's now hosted SPAs work :)
	muxrouter.NotFoundHandler = http.HandlerFunc(routes.spaIndexRoute)

	// Start server
	fmt.Printf("### 🌐 HTTP server listening on %v\n", serverPort)
	fmt.Printf("### 📁 Serving static content from '%v'\n", contentDir)
	err := http.ListenAndServe(":"+serverPort, muxrouter)
	if err != nil {
		log.Fatalln("Can't start server, that's super bad")
	}
}
