package main

import (
	"github.com/KrizzMU/coolback-alkol/internal/config"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/handler"
	"log"
)

func main() {
	cfg := config.MustLoad()

	// Setup REST server
	h := handler.New()
	s := rest.New(cfg.Server, h.InitRoutes())

	// TODO: Implement graceful shutdown
	if err := s.Run(); err != nil {
		log.Fatal("ERROR start Server!")
	}
}
