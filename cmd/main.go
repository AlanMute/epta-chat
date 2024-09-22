package main

import (
	"context"
	"fmt"
	"github.com/KrizzMU/coolback-alkol/internal/config"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	// Setup REST server
	h := handler.New()
	s := rest.New(cfg.Server, h.InitRoutes())

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Run(); err != nil {
			log.Printf(fmt.Sprintf("failed to start server because: %v", err))
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GracefulShutdownTimeout)*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Printf(fmt.Sprintf("failed to shutdown server because: %v", err))
		return
	}

	log.Print("stopped server")
}
