package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var contentDir string
var serverPort string

func main() {

	// Get server PORT setting or default
	serverPort = os.Getenv("PORT")
	if len(serverPort) == 0 {
		serverPort = "4000"
	}

	// Get CONTENT_DIR setting for static content or default
	contentDir = os.Getenv("CONTENT_DIR")
	if len(contentDir) == 0 && len(os.Args) > 1 {
		contentDir = os.Args[1]
	} else if len(contentDir) == 0 {
		contentDir = "."
	}

	// Routing
	muxrouter := mux.NewRouter()
	routes := Routes{
		contentDir:    contentDir,
		disableCORS:   true,
		darkskyAPIKey: "725f2b6bd8d8aa6ce91b85006771e89f",
		ipstackAPIKey: "e588291416844e390b0ea16b59671f30",
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

	// EVERYTHING else redirect to index.html
	muxrouter.NotFoundHandler = http.HandlerFunc(routes.spaIndexRoute) 

	// Start server
	fmt.Printf("### Starting server listening on %v\n", serverPort)
	fmt.Printf("### Serving static content from '%v'\n", contentDir)
	http.ListenAndServe(":"+serverPort, muxrouter)
}
