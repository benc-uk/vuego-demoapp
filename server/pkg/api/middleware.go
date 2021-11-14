package api

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"

	"github.com/elastic/go-sysinfo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Status struct {
	Hostname  string `json:"hostname"`
	Uptime    string `json:"uptime"`
	GoVersion string `json:"goVersion"`
	*Base
}

// Add logging middleware to the router in Apache Common Log Format.
func (b *Base) AddLogging(r *mux.Router) {
	r.Use(func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, next)
	})
}

// Add CORS middleware to the router
func (b *Base) AddCORS(origins []string, r *mux.Router) {
	r.Use(handlers.CORS(handlers.AllowedOrigins(origins)))
}

// AddMetrics adds Prometheus metrics to the router
func (b *Base) AddMetrics(r *mux.Router, prefix string) {
	r.Handle(prefix+"/metrics", promhttp.Handler())

	durationHistogram := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "response_duration_seconds",
		Help:        "A histogram of request latencies.",
		Buckets:     []float64{.001, .01, .1, .2, .5, 1, 2, 5},
		ConstLabels: prometheus.Labels{"handler": b.Name},
	}, []string{"method"})

	r.Use(func(next http.Handler) http.Handler {
		return promhttp.InstrumentHandlerDuration(durationHistogram, next)
	})
}

// AddHealth adds a health check endpoint to the API
func (b *Base) AddHealth(r *mux.Router, prefix string) {
	// Add health check endpoint
	r.HandleFunc(prefix+"/health", func(w http.ResponseWriter, r *http.Request) {
		if b.Healthy {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Service is not healthy"))
		}
	})
}

// AddStatus adds a status & info endpoint to the API
func (b *Base) AddStatus(r *mux.Router, prefix string) {
	r.HandleFunc(prefix+"/status", func(w http.ResponseWriter, r *http.Request) {
		host, _ := sysinfo.Host()
		host.Info().Uptime()
		status := Status{
			Hostname:  host.Info().Hostname,
			Uptime:    host.Info().Uptime().String(),
			GoVersion: runtime.Version(),
			Base:      b,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(status)
	})
}
