package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
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

func apiInfoRoute(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")

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

	/*if info.OS == "linux" {
		memtot, err := exec.Command("bash" ,"-c", "cat /proc/meminfo |grep MemTotal|grep -oP '\\d+'").Output()
		if err == nil {
			info.MemSystem, _ = strconv.ParseUint(strings.TrimSpace(string(memtot)), 10, 64)
		}
	}*/

	js, err := json.Marshal(info)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js)
}

func spaIndexRoute(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(resp, req, contentDir+"/index.html")
}
