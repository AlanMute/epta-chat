package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Address      string `yaml:"address" default:"0.0.0.0:8080"`
	ReadTimeout  uint   `yaml:"read_timeout" default:"10"`
	WriteTimeout uint   `yaml:"write_timeout" default:"10"`
	IdleTimeout  uint   `yaml:"idle_timeout" default:"10"`
}

type Server struct {
	httpServer *http.Server
}

func New(cfg Config, ginEngine *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           cfg.Address,
			Handler:        ginEngine,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
			IdleTimeout:    time.Duration(cfg.IdleTimeout) * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
