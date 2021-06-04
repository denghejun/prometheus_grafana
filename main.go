package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "denghejuncom",
		Name:      "myapp_processed_ops_total",
		Help:      "The total number of processed events",
	})
)

func main() {
	recordMetrics()
	startMetricsServer()
}

func startMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Print("prometheus server has been started at http://localhost:2112/metrics")
	http.ListenAndServe(":2112", nil) // http://localhost:2112
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
