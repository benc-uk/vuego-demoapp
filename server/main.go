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
	muxrouter.HandleFunc("/api/info", routes.apiInfoRoute)
	muxrouter.HandleFunc("/api/metrics", routes.apiMetricsRoute)
	muxrouter.HandleFunc("/api/weather", routes.weatherRoute)

	// These are SPA routes we want to handle in app, so redirect to index.html
	muxrouter.PathPrefix("/app").HandlerFunc(routes.spaIndexRoute)

	// Handle static content
	fileServer := http.FileServer(http.Dir(contentDir))
	muxrouter.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	// Start server
	fmt.Printf("### Starting server listening on %v\n", serverPort)
	fmt.Printf("### Serving static content from '%v'\n", contentDir)
	http.ListenAndServe(":"+serverPort, muxrouter)
}
