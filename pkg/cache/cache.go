package cache

import (
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/config"
	"github.com/vgekko/anistream-content/pkg/logger/sl"
	"log/slog"
)

func New(cfg config.Cache) *bigcache.BigCache {
	bc, err := bigcache.NewBigCache(bigcache.DefaultConfig(cfg.TTL))
	if err != nil {
		slog.Error("cannot initialize cache: ", sl.Err(err))
	}

	return bc
}
