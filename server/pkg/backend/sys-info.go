package backend

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
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

// Route to return system info
func (a *API) getInfo(resp http.ResponseWriter, req *http.Request) {
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

	if info.IsKubernetes {
		info.IsContainer = true
	}

	// Full grab of all env vars
	info.EnvVars = os.Environ()

	// Basic attempt to remove sensitive vars
	// Strange for-loop, means we can delete elements while looping over
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
	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write(js)
}

// Util to remove an element from a slice
func sliceRemove(slice []string, i int) []string {
	if i < len(slice)-1 {
		slice = append(slice[:i], slice[i+1:]...)
	} else if i == len(slice)-1 {
		slice = slice[:len(slice)-1]
	}
	return slice
}

// fileExists checks if a file or directory exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info != nil
}
