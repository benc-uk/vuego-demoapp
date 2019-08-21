package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/benc-uk/go-starter/pkg/envhelper"
)

var contentDir string
var serverPort string

func main() {
	// Get server PORT setting or default
	serverPort := envhelper.GetEnvString("PORT", "4000")
	// Get CONTENT_DIR setting for static content or default
	contentDir = envhelper.GetEnvString("CONTENT_DIR", ".") 

	// API keys
	darkskyAPIKey := envhelper.GetEnvString("WEATHER_API_KEY", "") 
	ipstackAPIKey := envhelper.GetEnvString("IPSTACK_API_KEY", "") 

	if(len(ipstackAPIKey) > 0) {
		fmt.Println("### Weather API enabled with DarkSky API key")
	}
	if(len(ipstackAPIKey) > 0) {
		fmt.Println("### Weather API enabled with IPStack API key")
	}

	// Routing
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

	// Handle static content, we have to explicitly put our top level dirs in here
	// - otherwise the NotFoundHandler will catch them
	fileServer := http.FileServer(http.Dir(contentDir))
	muxrouter.PathPrefix("/js").Handler(http.StripPrefix("/", fileServer))
	muxrouter.PathPrefix("/css").Handler(http.StripPrefix("/", fileServer))
	muxrouter.PathPrefix("/img").Handler(http.StripPrefix("/", fileServer))
	muxrouter.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/", fileServer))
	
	// EVERYTHING else redirect to index.html
	muxrouter.NotFoundHandler = http.HandlerFunc(routes.spaIndexRoute)

	// Start server
	fmt.Printf("### Starting server listening on %v\n", serverPort)
	fmt.Printf("### Serving static content from '%v'\n", contentDir)
	http.ListenAndServe(":"+serverPort, muxrouter)
}
