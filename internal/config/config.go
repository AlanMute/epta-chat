package config

import (
	"log"
	"os"

	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/pkg/logger/sl"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                     string      `yaml:"env" default:"local"`
	GracefulShutdownTimeout uint        `yaml:"graceful_shutdown_timeout" default:"10"`
	Server                  rest.Config `yaml:"server"`
	Logger                  sl.Config   `yaml:"logger"`
}

func MustLoad() *Config {
	configPath := "./config.yaml"

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	config := new(Config)

	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return config
}
