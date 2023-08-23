package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env   string `yaml:"env" env-default:"local"`
	GRPC  `yaml:"grpc"`
	Cache `yaml:"cache"`
}

type GRPC struct {
	Port string `yaml:"port" env-default:"50051"`
}

type Cache struct {
	TTL time.Duration `yaml:"ttl" env-default:"86400s"`
}

func Load() *Config {
	configPath := os.Getenv("CONTENT_SERVICE_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONTENT_SERVICE_CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
