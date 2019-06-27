package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// SysInfo is generic holder for passsing data back
type SysInfo struct {
	Hostname      string   `json:"hostname"`
	OS            string   `json:"os"`
	Arch          string   `json:"architecture"`
	CPUs          int      `json:"cpuCount"`
	GoVersion     string   `json:"goVersion"`
	NetRemoteAddr string   `json:"netRemoteAddress"`
	NetHost       string   `json:"netHost"`
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
// /api/info - Return system information and properties
//
func (r Routes) apiInfoRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}

	var info SysInfo

	// Grab various bits of infomation from where we can
	info.Hostname, _ = os.Hostname()
	info.GoVersion = runtime.Version()
	info.OS = runtime.GOOS
	info.Arch = runtime.GOARCH
	info.CPUs = runtime.NumCPU()
	info.NetRemoteAddr = req.RemoteAddr
	info.NetHost = req.Host
	info.EnvVars = os.Environ()

	// JSON-ify our info
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fire JSON result back down the internet tubes
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}

//
// /api/metrics - Return system metrics cpu, mem, etc
//
func (r Routes) apiMetricsRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}

	var metrics Metrics

	// Memory stuff
	memStats, err := mem.VirtualMemory()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.MemTotal = memStats.Total
	metrics.MemUsed = memStats.Used

	// CPU / processor stuff
	cpuStats, err := cpu.Percent(time.Millisecond*1000, false)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.CPUPerc = cpuStats[0]

	// Disk and filesystem usage stuff
	diskStats, err := disk.Usage("/")
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.DiskTotal = diskStats.Total
	metrics.DiskFree = diskStats.Free

	// Network stuff
	netStats, err := net.IOCounters(false)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.NetBytesRecv = netStats[0].BytesRecv
	metrics.NetBytesSent = netStats[0].BytesSent

	// JSON-ify our metrics
	js, err := json.Marshal(metrics)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fire JSON result back down the internet tubes
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}

//
// Special route to handle serving static SPA content with a JS router
//
func (r Routes) spaIndexRoute(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, contentDir + "/index.html")
}

//
// Weather info
//
func (r Routes) weatherRoute(resp http.ResponseWriter, req *http.Request) {
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}

	// Top level JSON container object
	var weather Weather

	// Try to guess calling IP address
	ip := req.Header.Get("x-forwarded-for")
	//ip = "86.134.117.146"
	if len(ip) == 0 {
		ip = req.RemoteAddr
	}
	if strings.HasPrefix(ip, "127.0.0.1") || strings.HasPrefix(ip, "[::1]") {
		http.Error(resp, "Localhost not allowed", http.StatusNotAcceptable)
		return
	}
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
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(apiresponse.Body)

	// Handle response and create object from JSON, and store in weather object
	var ipstackData ipstackAPIData
	err = json.Unmarshal(body, &ipstackData)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(ipstackData.City) == 0 {
		http.Error(resp, fmt.Sprintf("{\"msg\": \"Error no data for IP %v\"}", ip), http.StatusInternalServerError)
		return
	}
	weather.GeoInfo = ipstackData

	// Second API call is to Dark Sky to fetch weather data
	url = fmt.Sprintf("https://api.darksky.net/forecast/%s/%v,%v?exclude=minutely,hourly,daily&units=si", r.darkskyAPIKey, weather.GeoInfo.Lat, weather.GeoInfo.Long)
	apiresponse, err = netClient.Get(url)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err = ioutil.ReadAll(apiresponse.Body)

	// Handle response and create object from JSON, and store in weather object
	var darkskyData darkskyAPIData
	err = json.Unmarshal(body, &darkskyData)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	weather.WeatherInfo = darkskyData

	// JSON-ify our completed weather info object
	jsonResp, err := json.Marshal(weather)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	// Fire JSON result back down the internet tubes
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonResp)
}
