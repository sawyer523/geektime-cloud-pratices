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
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "code"})
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

func (t *ExecutionTimer) ObserveTotal(name string) {
	t.histo.WithLabelValues(name, "total").Observe(time.Now().Sub(t.start).Seconds())
}

func (t *ExecutionTimer) Count(name, code string) {
	t.count.WithLabelValues(name, code).Inc()
}

func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
}

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	count *prometheus.CounterVec
	start time.Time
	last  time.Time
}
