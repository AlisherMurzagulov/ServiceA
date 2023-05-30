package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"serviceB/cfg"
	"serviceB/internal/handlers"
	"serviceB/internal/services/rateLimiter"
	"serviceB/internal/services/serviceB"
	"serviceB/kafka"
	"serviceB/redis"
	"strconv"
	"time"
)

type Server interface {
	Start() (serverChannel chan error)
	Stop() error
}

type server struct {
	router        *gin.Engine
	srv           *http.Server
	srvCh         chan error
	stopTimeoutMS time.Duration
	handlers      handlers.Manager
}

func (s *server) Start() (serverChannel chan error) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.srvCh <- err
		}
	}()

	return s.srvCh
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.stopTimeoutMS)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server forced to shutdown: %v", err)
	}
	return nil
}

func NewServer(redis redis.Manager, kafka kafka.Manager) (s Server, err error) {
	if cfg.AppConfigs.WebServer.Port <= 0 {
		return nil, fmt.Errorf("bad port value %d for server", cfg.AppConfigs.WebServer.Port)
	}

	if cfg.AppConfigs.StopTimeoutMS < 1 {
		return nil, fmt.Errorf("bad stop timeout %d for server", cfg.AppConfigs.StopTimeoutMS)
	}

	if err != nil {
		return nil, err
	}

	router := router(cfg.AppConfigs.WebServer.GIN.ReleaseMode, cfg.AppConfigs.WebServer.GIN.UseLogger, cfg.AppConfigs.WebServer.GIN.UseRecovery)
	limiter := rateLimiter.NewRateLimiter(redis)
	serviceB := serviceB.NewServiceB(redis, kafka, limiter)

	srv := &server{
		srv: &http.Server{
			Addr:    ":" + strconv.Itoa(cfg.AppConfigs.WebServer.Port),
			Handler: router,
		},
		router:        router,
		stopTimeoutMS: time.Duration(cfg.AppConfigs.StopTimeoutMS) * time.Millisecond,
		handlers:      handlers.NewManager(serviceB),
	}
	srv.initEndpoints()

	return srv, nil
}
