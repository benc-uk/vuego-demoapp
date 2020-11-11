package main

import (
	"fmt"
	"net/http"

	"github.com/benc-uk/go-starter/pkg/envhelper"
	"github.com/gorilla/mux"
)

var contentDir string
var serverPort string

func redirectToHttps(w http.ResponseWriter, r *http.Request){
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
        target += "?" + r.URL.RawQuery
    }
	//fmt.Println("Redirecting to: ", target)
	http.Redirect(w, r, target, http.StatusTemporaryRedirect)	//TemporaryRedirect im important so that the POST body comes with!
}

func main() {
	// Get server PORT setting or default
	serverPort := envhelper.GetEnvString("PORT", "443")
	// Get CONTENT_DIR setting for static content or default
	contentDir = envhelper.GetEnvString("CONTENT_DIR", "/var/www/vuego-demoapp/spa/dist/")		//important to "npm run build" inside /spa folder

	// API keys
	darkskyAPIKey := envhelper.GetEnvString("WEATHER_API_KEY", "")
	ipstackAPIKey := envhelper.GetEnvString("IPSTACK_API_KEY", "")

	if len(darkskyAPIKey) > 0 {
		fmt.Println("### Weather API enabled with DarkSky API key")
	}
	if len(ipstackAPIKey) > 0 {
		fmt.Println("### Weather API enabled with IPStack API key")
	}

	// Certificate paths
	fullchainPath := "/etc/letsencrypt/live/[ENTER DOMAIN NAME]/fullchain.pem"
	privkeyPath := "/etc/letsencrypt/live/[ENTER DOMAIN NAME]/privkey.pem"

	if (fullchainPath == "/etc/letsencrypt/live/[ENTER DOMAIN NAME]/fullchain.pem") {
		fmt.Println("### Path to your private certificate is not defined, LetsEncrypt uses /etc/letsencrypt/live/[ENTER DOMAIN NAME]/fullchain.pem by default")
	}
	if (privkeyPath == "/etc/letsencrypt/live/[ENTER DOMAIN NAME]/privkey.pem") {
		fmt.Println("### Path to your private key is not defined, LetsEncrypt uses /etc/letsencrypt/live/[ENTER DOMAIN NAME]/privkey.pem by default")
	}

	// Routing
	go http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
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
	http.ListenAndServeTLS(":"+serverPort, fullchainPath, privkeyPath, muxrouter)
}
