package main

import (
	"context"
	"fmt"
	"github.com/KrizzMU/coolback-alkol/internal/config"
	"github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/model"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/handler"
	"github.com/KrizzMU/coolback-alkol/pkg/logger/sl"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env, cfg.Logger)

	log.With("config", cfg).Info("Application start!")

	messenger := model.NewMessenger()

	// Setup REST server
	h := handler.New(messenger)
	s := rest.New(cfg.Server, h.InitRoutes())

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Run(); err != nil {
			log.Error(fmt.Sprintf("failed to start server because: %v", err))
		}
	}()

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GracefulShutdownTimeout)*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("failed to shutdown server because: %v", err))
		return
	}

	log.Info("Application stopped!")
}
