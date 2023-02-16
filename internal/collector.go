package internal

import "github.com/prometheus/client_golang/prometheus"

var (
	LatencyDistributionCollector = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "ws",
		Subsystem: "request",
		Name:      "latency_distribution",
		Help:      "μs",
		Buckets:   []float64{5, 10, 50, 100, 500, 1000},
	}, []string{"application"})

	LatencyPercentileCollector = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "ws",
		Subsystem: "request",
		Name:      "latency_percentile",
		Help:      "",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	}, []string{"application"})
)

func init() {
	_ = prometheus.Register(LatencyDistributionCollector)
	_ = prometheus.Register(LatencyPercentileCollector)
}
