package dbConfig

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func GetConnectionString() string {
	configData, err := os.ReadFile("./internal/config/dbConf/conf.yaml")
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.User, config.Password,
		config.Dbname, config.Host, config.Port)

	logrus.Info(connectionString)

	return connectionString
}
