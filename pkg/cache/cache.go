package cache

import (
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/config"
	"log/slog"
)

func New(cfg config.Cache) *bigcache.BigCache {
	bc, err := bigcache.NewBigCache(bigcache.DefaultConfig(cfg.TTL))
	if err != nil {
		slog.Error("failed to initialize bigcache")
		return nil
	}

	return bc
}
