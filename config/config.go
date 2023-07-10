package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	Postgres `yaml:"postgres"`
	Redis `yaml:"redis"`
}

type Postgres struct {
	PoolMax int `yaml:"pool_max" env-default:"2"`
	URL string `yaml:"url" env-required:"true"`
}

type Redis struct {
	InfoCacheTTL time.Duration ` yaml:"info_cache_ttl" env-required:"true"`
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

