package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env   string `yaml:"env" env-default:"local"`
	HTTP  `yaml:"http"`
	GRPC  `yaml:"grpc"`
	Redis `yaml:"redis"`
}

type HTTP struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8800"`
}

type GRPC struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"50051"`
}

type Redis struct {
	Addr         string        `yaml:"addr" env-required:"true"`
	InfoCacheTTL time.Duration `yaml:"info_cache_ttl" env-default:"86400s"`
}

func Load() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
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
