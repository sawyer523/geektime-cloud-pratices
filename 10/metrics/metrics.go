package metrics

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func Register() {
	err := prometheus.Register(_metricSeconds)
	if err != nil {
		fmt.Println(err)
	}

	err = prometheus.Register(_metricRequests)
	if err != nil {
		fmt.Println(err)
	}
}

var (
	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "httpserver",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"func", "method", "kind"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "httpserver",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"func", "method", "code"})
)

func New() *ExecutionTimer {
	return NewExecutionTimer(_metricSeconds, _metricRequests)
}

func NewExecutionTimer(histo *prometheus.HistogramVec, count *prometheus.CounterVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: histo,
		count: count,
		start: now,
		last:  now,
	}
}

func (t *ExecutionTimer) ObserveTotal(name, method string) {
	t.histo.WithLabelValues(name, method, "total").Observe(time.Now().Sub(t.start).Seconds())
}

func (t *ExecutionTimer) Count(name, method, code string) {
	t.count.WithLabelValues(name, method, code).Inc()
}

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	count *prometheus.CounterVec
	start time.Time
	last  time.Time
}
