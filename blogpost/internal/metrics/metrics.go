package metrics

import (
	apimetrics "blogpost/internal/api/metrics"
)

type Metrics struct {
	HTTPAPI *apimetrics.Metrics
}

func NewMetrics() *Metrics {
	return &Metrics{
		HTTPAPI: apimetrics.NewMetrics(),
	}
}
