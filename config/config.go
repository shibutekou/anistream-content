package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		PG PG `yaml:"postgres"`
		Redis Redis `yaml:"redis"`
	}

	PG struct {
		PoolMax int `env-required:"true" env:"PG_POOL_MAX" yaml:"pool_max"`
		URL string `env-required:"true" env:"PG_URL"`
	}

	Redis struct {
		InfoCacheTTL int `env-required:"true" env:"INFO_CACHE_TTL" yaml:"info_cache_ttl"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config.NewConfig: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

