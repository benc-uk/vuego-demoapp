package backend

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// Force garbage collection
func (a *API) getRunGC(resp http.ResponseWriter, req *http.Request) {
	runtime.GC()

	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write([]byte("Garbage collector was run"))
}

// Allocate a lot of memory
func (a *API) postAllocMem(resp http.ResponseWriter, req *http.Request) {
	params := struct {
		Size int
	}{
		Size: 100,
	}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	if params.Size < 1 {
		params.Size = 1
	}
	if params.Size > 400 {
		params.Size = 400
	}

	var buffer = make([]string, params.Size*1024*1024)
	for e := range buffer {
		buffer[e] = "#"
	}

	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write([]byte("Memory was allocated"))
}

// Force max CPU load for a number of seconds
func (a *API) postForceCPU(resp http.ResponseWriter, req *http.Request) {
	params := struct {
		Seconds int
	}{
		Seconds: 1,
	}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	if params.Seconds < 1 {
		params.Seconds = 1
	}
	if params.Seconds > 10 {
		params.Seconds = 10
	}

	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	quit := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-quit:
					return
				default: //nolint
				}
			}
		}()
	}

	time.Sleep(time.Duration(params.Seconds) * time.Second)
	for i := 0; i < n; i++ {
		quit <- true
	}

	resp.Header().Set("Content-Type", "application/json")
	_, _ = resp.Write([]byte("CPU load was forced"))
}
