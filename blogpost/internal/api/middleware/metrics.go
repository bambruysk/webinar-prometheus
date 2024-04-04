package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Metrics interface {
	IncBlogPostCreate()
	IncBlogPostGet()
	IncBlogPostGetError(status int)
	IncBlogPostCreateError(status int)
	IncBlogPostHistogram(method, path string, status int, duration float64)
}

type MetricsMiddleware struct {
	m Metrics
}

func NewRequestDurationMiddleware(metrics Metrics) *MetricsMiddleware {
	return &MetricsMiddleware{
		m: metrics,
	}
}

func (m *MetricsMiddleware) Duration(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		start := time.Now()
		err := next(ectx)
		duration := time.Since(start).Seconds()
		m.m.IncBlogPostHistogram(ectx.Request().Method, ectx.Path(), ectx.Response().Status, duration)

		return err
	}
}

func (m *MetricsMiddleware) GetError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		err := next(ectx)
		if err != nil {
			m.m.IncBlogPostGetError(ectx.Response().Status)
		}

		return err
	}
}

func (m *MetricsMiddleware) CreateError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		err := next(ectx)
		if err != nil {
			m.m.IncBlogPostCreateError(ectx.Response().Status)
		}

		return err
	}
}

func (m *MetricsMiddleware) TotalGet(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		m.m.IncBlogPostGet()
		return next(ectx)
	}
}

func (m *MetricsMiddleware) TotalCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		m.m.IncBlogPostCreate()
		return next(ectx)
	}
}
