package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// SysInfo is generic holder for passsing data back
type SysInfo struct {
	Hostname      string
	OS            string
	Arch          string
	Cpus          int
	GoVersion     string
	NetRemoteAddr string
	NetHost       string
}

// Real time system metrics
type Metrics struct {
	MemPercent float64
	MemTotal   uint64
	MemUsed    uint64
	Cpu        float64
}

func apiInfoRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	resp.Header().Set("Access-Control-Allow-Origin", "*")

	var info SysInfo

	// Grab various bits of infomation from where we can
	info.Hostname, _ = os.Hostname()
	info.GoVersion = runtime.Version()
	info.OS = runtime.GOOS
	info.Arch = runtime.GOARCH
	info.Cpus = runtime.NumCPU()
	info.NetRemoteAddr = req.RemoteAddr
	info.NetHost = req.Host

	// JSON-ify our info
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	// Shoot JSON result back down them internet tubes
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}

func apiMetricsRoute(resp http.ResponseWriter, req *http.Request) {
	// CORS is for wimps
	resp.Header().Set("Access-Control-Allow-Origin", "*")

	var metrics Metrics

	// Memory stuff
	memStats, err := mem.VirtualMemory()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.MemPercent = memStats.UsedPercent
	metrics.MemTotal = memStats.Total
	metrics.MemUsed = memStats.Used

	// CPU / processor stuff
	cpuStats, err := cpu.Percent(time.Millisecond*1000, false)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.Cpu = cpuStats[0]

	// JSON-ify our metrics
	js, err := json.Marshal(metrics)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	// Shoot JSON result back down them internet tubes
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}

func spaIndexRoute(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(resp, req, contentDir+"/index.html")
}
