package main

import (
	"context"
	"fmt"
	messenger_service "github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/config"
	"github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/model"
	"github.com/KrizzMU/coolback-alkol/internal/db"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/handler"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"github.com/KrizzMU/coolback-alkol/pkg/logger/sl"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env, cfg.Logger)
	log.With("config", cfg).Info("Application start!")

	repositories := repository.New(db.GetConnection())
	if err := repositories.EnsureCommonChatExists(); err != nil {
		logrus.Fatal("Failed to creat common chat: ", err)
	}

	tokenManager, err := auth.NewManager("hello-world")
	if err != nil {
		panic(err)
	}

	services := service.New(repositories, tokenManager)

	messenger := model.NewMessenger()

	messengerService, err := messenger_service.NewMessenger(
		repositories.Chat,
		repositories.User,
		repositories.Message,
		messenger,
	)
	if err != nil {
		logrus.Fatal("Failed to creat common chat: ", err)
	}

	// Setup REST server
	h := handler.New(tokenManager, services, messengerService)
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
