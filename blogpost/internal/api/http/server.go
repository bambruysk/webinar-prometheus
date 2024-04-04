package http

import (
	"context"
	"errors"
	"log"
	"net/http"

	apimw "blogpost/internal/api/middleware"
	apimodels "blogpost/internal/api/models"

	"blogpost/internal/models"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/labstack/echo/v4"
)

type BlogPostService interface {
	Create(ctx context.Context, post *models.BlogPost) (*models.BlogPost, error)
	Get(ctx context.Context, id string) (*models.BlogPost, error)
}

type Metrics interface {
	IncBlogPostCreate()
	IncBlogPostGet()
	IncBlogPostGetError(status int)
	IncBlogPostCreateError(status int)
	IncBlogPostHistogram(method, path string, status int, duration float64)

	Register() prometheus.Registerer
}

type Server struct {
	e *echo.Echo

	metricServer *echo.Echo

	service BlogPostService
	metrics Metrics

	statsMiddleware *apimw.MetricsMiddleware
}

func NewServer(service BlogPostService, metrics Metrics) *Server {
	return &Server{
		e:               echo.New(),
		metricServer:    echo.New(),
		service:         service,
		metrics:         metrics,
		statsMiddleware: apimw.NewRequestDurationMiddleware(metrics),
	}
}

func (s *Server) Run(ctx context.Context) error {
	s.e.Use(s.statsMiddleware.Duration)
	s.e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		Registerer: s.metrics.Register(),
	}))

	s.e.POST("/posts", s.CreatePost, s.statsMiddleware.TotalCreate, s.statsMiddleware.CreateError)
	s.e.GET("/posts/:id", s.GetPost, s.statsMiddleware.TotalGet, s.statsMiddleware.GetError)

	s.metricServer.GET("/metrics", echoprometheus.NewHandler())

	go func() {
		if err := s.metricServer.Start(":18080"); err != nil {
			log.Println(err)
		}
	}()

	return s.e.Start(":8080")
}

func (s *Server) CreatePost(ectx echo.Context) error {
	post := &apimodels.CreateBlogPost{}
	if err := ectx.Bind(post); err != nil {
		return err
	}

	created, err := s.service.Create(ectx.Request().Context(), apimodels.ConvertToBlogPost(post))
	if err != nil {
		if errors.Is(err, models.ErrorAlreadyExists) {
			return ectx.String(http.StatusConflict, err.Error())
		}

		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.String(http.StatusCreated, created.ID)
}

func (s *Server) GetPost(ectx echo.Context) error {
	id := ectx.Param("id")
	post, err := s.service.Get(ectx.Request().Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrorPostNotFound) {
			return ectx.String(http.StatusNotFound, err.Error())
		}

		return ectx.String(http.StatusInternalServerError, err.Error())
	}

	return ectx.JSON(http.StatusOK, apimodels.ConvertToGetBlogPost(post))
}
