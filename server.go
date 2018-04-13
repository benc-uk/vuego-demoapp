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
	router := mux.NewRouter()
	router.HandleFunc("/api/info", apiInfoRoute)

	// These are SPA routes we want to handle in app, so redirect to index.html
	router.PathPrefix("/app").HandlerFunc(spaIndexRoute)

	// Handle static content
	fileServer := http.FileServer(http.Dir(contentDir))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	// Start server
	fmt.Printf("### Starting server listening on %v\n", serverPort)
	fmt.Printf("### Serving static content from '%v'\n", contentDir)
	http.ListenAndServe(":"+serverPort, router)
}
