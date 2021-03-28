package main

import (
	// #include <unistd.h>
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/pbnjay/memory"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// SysInfo is generic holder for passsing data back
type SysInfo struct {
	Hostname      string   `json:"hostname"`
	Platform      string   `json:"platform"`
	OS            string   `json:"os"`
	Uptime        uint64   `json:"uptime"`
	Arch          string   `json:"architecture"`
	CPUs          int      `json:"cpuCount"`
	CPUModel      string   `json:"cpuModel"`
	Mem           uint64   `json:"mem"`
	GoVersion     string   `json:"goVersion"`
	NetRemoteAddr string   `json:"netRemoteAddress"`
	NetHost       string   `json:"netHost"`
	IsContainer   bool     `json:"isContainer"`
	IsKubernetes  bool     `json:"isKubernetes"`
	EnvVars       []string `json:"envVars"`
}

// Metrics are real time system counters
type Metrics struct {
	MemTotal     uint64  `json:"memTotal"`
	MemUsed      uint64  `json:"memUsed"`
	CPUPerc      float64 `json:"cpuPerc"`
	DiskTotal    uint64  `json:"diskTotal"`
	DiskFree     uint64  `json:"diskFree"`
	NetBytesSent uint64  `json:"netBytesSent"`
	NetBytesRecv uint64  `json:"netBytesRecv"`
}

// Weather holds data about the weather
type Weather struct {
	IPAddress   string         `json:"ipAddress"`
	GeoInfo     ipstackAPIData `json:"location"`
	WeatherInfo darkskyAPIData `json:"weather"`
}

// HTTPError holds API JSON error
type HTTPError struct {
	Error string `json:"error"`
}

// ipstackAPIData holds results of IPStack API call
type ipstackAPIData struct {
	City    string  `json:"city"`
	Country string  `json:"country_name"`
	Lat     float64 `json:"latitude"`
	Long    float64 `json:"longitude"`
}

// darkskyAPIData holds results of Dark Sky API call
type darkskyAPIData struct {
	Timezone  string `json:"timezone"`
	Currently struct {
		Summary           string  `json:"summary"`
		Icon              string  `json:"icon"`
		PrecipProbability float32 `json:"precipProbability"`
		Temperature       float32 `json:"temperature"`
		WindSpeed         float32 `json:"windSpeed"`
		UVIndex           float32 `json:"uvIndex"`
		Humidity          float32 `json:"humidity"`
	} `json:"currently"`
}

// Routes is our exported class
type Routes struct {
	contentDir    string
	disableCORS   bool
	darkskyAPIKey string
	ipstackAPIKey string
}

//
// Simple config API for the frontend
//
func (r Routes) configRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}
	resp.Header().Set("Content-Type", "application/json")

	// Passes through a single env var, AUTH_CLIENT_ID
	config := struct {
		AuthClientAd string `json:"AUTH_CLIENT_ID"`
	}{}
	config.AuthClientAd = os.Getenv("AUTH_CLIENT_ID")

	jsonResp, err := json.Marshal(config)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	// Fire JSON result back down the internet tubes
	_, _ = resp.Write(jsonResp)
}

//
// /api/info - Return system information and properties
//
func (r Routes) apiInfoRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}
	resp.Header().Set("Content-Type", "application/json")

	var info SysInfo

	hostInfo, err := host.Info()
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	cpuInfo, err := cpu.Info()
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	// Grab various bits of infomation from where we can
	info.Hostname, _ = os.Hostname()
	info.GoVersion = runtime.Version()
	info.OS = hostInfo.Platform + " " + hostInfo.PlatformVersion
	info.Platform = hostInfo.OS
	info.Uptime = hostInfo.Uptime
	info.Mem = memory.TotalMemory()
	info.Arch = runtime.GOARCH
	info.CPUs = runtime.NumCPU()
	info.CPUModel = cpuInfo[0].ModelName
	info.NetRemoteAddr = req.RemoteAddr
	info.NetHost = req.Host
	info.IsContainer = fileExists("/.dockerenv")
	info.IsKubernetes = fileExists("/var/run/secrets/kubernetes.io")

	// Full grab of all env vars
	info.EnvVars = os.Environ()

	// Basic attempt to remove sensitive vars
	// Strange for means we can delete elements while looping over
	for i := len(info.EnvVars) - 1; i >= 0; i-- {
		envVarName := strings.Split(info.EnvVars[i], "=")[0]
		if strings.Contains(envVarName, "_KEY") || strings.Contains(envVarName, "SECRET") || strings.Contains(envVarName, "PWD") || strings.Contains(envVarName, "PASSWORD") {
			info.EnvVars = sliceRemove(info.EnvVars, i)
		}
	}

	// JSON-ify our info
	js, err := json.Marshal(info)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	// Fire JSON result back down the internet tubes
	_, _ = resp.Write(js)
}

//
// /api/metrics - Return system metrics cpu, mem, etc
//
func (r Routes) apiMetricsRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}
	resp.Header().Set("Content-Type", "application/json")

	var metrics Metrics

	// Memory stuff
	memStats, err := mem.VirtualMemory()
	if err != nil {
		apiError(resp, http.StatusInternalServerError, "Virtual memory "+err.Error())
		return
	}
	metrics.MemTotal = memStats.Total
	metrics.MemUsed = memStats.Used

	// CPU / processor stuff
	cpuStats, err := cpu.Percent(0, false)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, "CPU percentage "+err.Error())
		return
	}
	metrics.CPUPerc = cpuStats[0]

	// Disk and filesystem usage stuff
	diskStats, err := disk.Usage("/")
	if err != nil {
		apiError(resp, http.StatusInternalServerError, "Disk usage "+err.Error())
		return
	}
	metrics.DiskTotal = diskStats.Total
	metrics.DiskFree = diskStats.Free

	// Network stuff
	netStats, err := net.IOCounters(false)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, "IOCounters "+err.Error())
		return
	}
	metrics.NetBytesRecv = netStats[0].BytesRecv
	metrics.NetBytesSent = netStats[0].BytesSent

	// JSON-ify our metrics
	js, err := json.Marshal(metrics)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	// Fire JSON result back down the internet tubes
	_, _ = resp.Write(js)
}

//
// Special route to handle serving static SPA content with a JS router
//
func (r Routes) spaIndexRoute(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, contentDir+"/index.html")
}

//
// Weather info
//
func (r Routes) weatherRoute(resp http.ResponseWriter, req *http.Request) {
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}
	resp.Header().Set("Content-Type", "application/json")

	// Check if required config is set
	if r.ipstackAPIKey == "" {
		apiError(resp, http.StatusNotImplemented, "Feature disabled, IPSTACK_API_KEY is not set")
		return
	}
	if r.darkskyAPIKey == "" {
		apiError(resp, http.StatusNotImplemented, "Feature disabled, WEATHER_API_KEY is not set")
		return
	}

	// Top level JSON container struct
	var weather Weather

	// Try to deduce calling IP address
	ip := req.Header.Get("x-forwarded-for")

	// Special trick to work with local dev
	if hostname, _ := os.Hostname(); hostname == "BENSL3" {
		ip = "212.36.160.18" // Only uncomment for local testing!
	}

	// If not in the header try the RemoteAddr field
	if len(ip) == 0 {
		ip = req.RemoteAddr
	}

	// Checks for localhost, as it won't work
	if strings.HasPrefix(ip, "127.0.0.1") || strings.HasPrefix(ip, "[::1]") {
		apiError(resp, http.StatusNotAcceptable, fmt.Sprintf("This IP is not allowed %v", ip))
		return
	}

	// There might be a port in there, get rid of it
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}

	weather.IPAddress = ip

	// First API call is to IPStack to reverse lookup IP into location (lat & long)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var netClient = &http.Client{Timeout: time.Second * 10, Transport: tr}

	url := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s&format=1", ip, r.ipstackAPIKey)
	apiresponse, err := netClient.Get(url)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	body, _ := ioutil.ReadAll(apiresponse.Body)

	// Handle response and create object from JSON, and store in weather object
	var ipstackData ipstackAPIData
	err = json.Unmarshal(body, &ipstackData)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	if len(ipstackData.City) == 0 {
		apiError(resp, http.StatusNotFound, fmt.Sprintf("No location data for this IP %v", ip))
		return
	}
	weather.GeoInfo = ipstackData

	// Second API call is to Dark Sky to fetch weather data
	url = fmt.Sprintf("https://api.darksky.net/forecast/%s/%v,%v?exclude=minutely,hourly,daily&units=si", r.darkskyAPIKey, weather.GeoInfo.Lat, weather.GeoInfo.Long)
	apiresponse, err = netClient.Get(url)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	body, _ = ioutil.ReadAll(apiresponse.Body)

	// Handle response and create object from JSON, and store in weather object
	var darkskyData darkskyAPIData
	err = json.Unmarshal(body, &darkskyData)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	// Push API response data into our weather struct
	weather.WeatherInfo = darkskyData

	// JSON-ify our completed weather info object
	jsonResp, err := json.Marshal(weather)
	if err != nil {
		apiError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	// Fire JSON result back down the internet tubes
	_, _ = resp.Write(jsonResp)
}

//
// Helper function for returning API errors
//
func apiError(resp http.ResponseWriter, code int, message string) {
	resp.WriteHeader(code)

	//message = strings.ReplaceAll(message, "\"", "'")
	errorData := &HTTPError{
		Error: message,
	}

	errorResp, err := json.Marshal(errorData)
	if err != nil {
		fmt.Printf("### ERROR! httpError unable to marshal to JSON. Message was %s\n", message)
		return
	}
	_, _ = resp.Write(errorResp)
}

// fileExists checks if a file or directory exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	//return !info.IsDir()
	return info != nil
}

func sliceRemove(slice []string, i int) []string {
	if i < len(slice)-1 {
		slice = append(slice[:i], slice[i+1:]...)
	} else if i == len(slice)-1 {
		slice = slice[:len(slice)-1]
	}
	return slice
}
