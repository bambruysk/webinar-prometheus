package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	CreateCounter prometheus.Counter
	GetCounter    prometheus.Counter
	CreateError   *prometheus.CounterVec
	GetError      *prometheus.CounterVec

	BlogHistogram *prometheus.HistogramVec

	register prometheus.Registerer
}

func NewMetrics() *Metrics {
	m := &Metrics{
		CreateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "create_total",
		}),
		GetCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "get_total",
		}),
		CreateError: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "create_error_total",
		}, []string{"status"}),
		GetError: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "get_error_total",
		}, []string{"status"}),
		BlogHistogram: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "response_time_seconds",
			Help: "BlogPost histogram",
		}, []string{"method", "path", "status"}),
	}

	m.register = prometheus.WrapRegistererWithPrefix("blogpost_http_", prometheus.DefaultRegisterer)

	m.register.MustRegister(
		m.CreateCounter,
		m.GetCounter,
		m.CreateError,
		m.GetError,
		m.BlogHistogram,
	)

	return m
}

func (m *Metrics) IncBlogPostCreate() {
	m.CreateCounter.Inc()
}

func (m *Metrics) IncBlogPostGet() {
	m.GetCounter.Inc()
}

func (m *Metrics) IncBlogPostGetError(status int) {
	m.CreateError.With(prometheus.Labels{"status": strconv.Itoa(status)}).Inc()
}

func (m *Metrics) IncBlogPostCreateError(status int) {
	m.CreateError.With(prometheus.Labels{"status": strconv.Itoa(status)}).Inc()
}

func (m *Metrics) IncBlogPostHistogram(method, path string, status int, duration float64) {
	m.BlogHistogram.With(prometheus.Labels{"method": method, "path": path, "status": strconv.Itoa(status)}).Observe(duration)
}

func (m *Metrics) Register() prometheus.Registerer {
	return m.register
}
