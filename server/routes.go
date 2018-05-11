package main

import (
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
	Cpus          int      `json:"cpuCount"`
	GoVersion     string   `json:"goVersion"`
	NetRemoteAddr string   `json:"netRemoteAddress"`
	NetHost       string   `json:"netHost"`
	EnvVars       []string `json:"envVars"`
}

// Real time system metrics
type Metrics struct {
	MemTotal     uint64  `json:"memTotal"`
	MemUsed      uint64  `json:"memUsed"`
	CpuPerc      float64 `json:"cpuPerc"`
	DiskTotal    uint64  `json:"diskTotal"`
	DiskFree     uint64  `json:"diskFree"`
	NetBytesSent uint64  `json:"netBytesSent"`
	NetBytesRecv uint64  `json:"netBytesRecv"`
}

// Real time system metrics
type Weather struct {
	IpAddress string         `json:"ipAddress"`
	GeoInfo   ipstackApiData `json:"geo"`
}

// Real time system metrics
type ipstackApiData struct {
	City    string  `json:"city"`
	Country string  `json:"country_name"`
	Lat     float64 `json:"latitude"`
	Long    float64 `json:"longitude"`
}

type Routes struct {
	contentDir    string
	disableCORS   bool
	darkskyApiKey string
	ipstackApiKey string
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
	info.Cpus = runtime.NumCPU()
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
	metrics.CpuPerc = cpuStats[0]

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
// Special route to handle serving
//
func (r Routes) spaIndexRoute(resp http.ResponseWriter, req *http.Request) {
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}

	http.ServeFile(resp, req, contentDir+"/index.html")
}

//
// Weather info
//
func (r Routes) weatherRoute(resp http.ResponseWriter, req *http.Request) {
	if r.disableCORS {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
	}

	var weather Weather

	ip := req.Header.Get("x-forwarded-for")
	if len(ip) == 0 {
		ip = req.RemoteAddr
	}
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]

	}
	weather.IpAddress = ip

	var netClient = &http.Client{Timeout: time.Second * 10}
	url := fmt.Sprintf("http://api.ipstack.com/%v?access_key=%v", ip, r.ipstackApiKey)
	apiresponse, err := netClient.Get(url)
	body, err := ioutil.ReadAll(apiresponse.Body)
	var ipstackData ipstackApiData
	err = json.Unmarshal(body, &ipstackData)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(ipstackData.City) == 0 {
		http.Error(resp, string(body), http.StatusInternalServerError)
		return
	}
	weather.GeoInfo = ipstackData

	// JSON-ify our weather info
	js, err := json.Marshal(weather)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	// Fire JSON result back down the internet tubes
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}
