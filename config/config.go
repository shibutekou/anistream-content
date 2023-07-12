package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env string 	`yaml:"env" env-default:"local"`
	HTTP 		`yaml:"http"`
	Postgres 	`yaml:"postgres"`
	Redis 		`yaml:"redis"`
}

type HTTP struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8800"`
}

type Postgres struct {
	MaxPoolSize int 			`yaml:"max_pool_size" env-default:"2"`
	ConnAttempts int 			`yaml:"conn_attempts" env_default:"10"`
	ConnTimeout time.Duration 	`yaml:"conn_timeout" env-default:"1s"`

	Username string 			`yaml:"username" env-required:"true"`
	Host string 				`yaml:"host" env-required:"true"`
	Port string 				`yaml:"port" env-required:"true"`
	Database string				`yaml:"database" env-required:"true"`
	Password string      		`yaml:"password" env-required:"true"`
	SSLMode string 				`yaml:"ssl_mode" env-default:"disable"`
}

type Redis struct {
	Addr string 				`yaml:"addr" env-required:"true"`
	InfoCacheTTL time.Duration 	`yaml:"info_cache_ttl" env-required:"true"`
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

