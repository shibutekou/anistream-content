package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vgekko/ani-go/internal/entity"
)

const notExists = iota

type InfoRepo struct {
	db  *redis.Client
	ttl time.Duration
}

func NewInfoRepo(db *redis.Client, ttl time.Duration) *InfoRepo {
	return &InfoRepo{db, ttl}
}

func (r *InfoRepo) Cache(ctx context.Context, key string, value []entity.TitleInfo) error {
	var b bytes.Buffer

	// Create an encoder and encode the value.
	enc := gob.NewEncoder(&b)
	_ = enc.Encode(value)

	// Get the bytes from the buffer
	byteSlice := b.Bytes()

	err := r.db.Set(ctx, key, byteSlice, r.ttl).Err()
	if err != nil {
		return fmt.Errorf("InfoRedisRepo.Cache: %w", err)
	}

	return nil
}

func (r *InfoRepo) Lookup(ctx context.Context, key string) bool {
	data, _ := r.db.Exists(ctx, key).Result()

	return data != notExists
}

func (r *InfoRepo) FromCache(ctx context.Context, key string) ([]entity.TitleInfo, error) {
	var ti = make([]entity.TitleInfo, 0)
	var asBytes []byte

	if err := r.db.Get(ctx, key).Scan(&asBytes); err != nil {
		return nil, fmt.Errorf("InfoRedisRepo.FromCache: %w", err)
	}

	b := bytes.NewReader(asBytes)

	// Create a decoder and decode the value
	_ = gob.NewDecoder(b).Decode(&ti)

	return ti, nil
}
