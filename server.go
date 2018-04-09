package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"os/exec"
	"strings"
	"strconv"

	"github.com/gorilla/mux"
)

// SysInfo is generic holder for passsing data to the templates
type SysInfo struct {
	Hostname      string
	OS            string
	Arch          string
	Cpus          int
	GoVersion     string
	NetRemoteAddr string
	NetHost       string
	MemAlloc      uint64
	MemSystem     uint64
}

var contentDir string
var serverPort string

func apiInfoRoute(resp http.ResponseWriter, req *http.Request) {

	var info SysInfo
	info.Hostname, _ = os.Hostname()
	info.GoVersion = runtime.Version()
	info.OS = runtime.GOOS
	info.Arch = runtime.GOARCH
	info.Cpus = runtime.NumCPU()
	info.NetRemoteAddr = req.RemoteAddr
	info.NetHost = req.Host

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	info.MemAlloc = mem.TotalAlloc

	if info.OS == "linux" { 
		memtot, err := exec.Command("bash" ,"-c", "cat /proc/meminfo |grep MemTotal|grep -oP '\\d+'").Output()
		if err == nil {
			info.MemSystem, _ = strconv.ParseUint(strings.TrimSpace(string(memtot)), 10, 64)
		}
	}

	js, err := json.Marshal(info)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}

func spaIndexRoute(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, contentDir+"/index.html")
}

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
