package backend

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

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

// Route to return system metrics cpu, mem, etc
// NOTE! This is not the same as /api/metrics used for Prometheus, registered with AddMetrics
func (a *API) getMonitorMetrics(resp http.ResponseWriter, req *http.Request) {
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
	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write(js)
}
