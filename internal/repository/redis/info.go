package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vgekko/anistream-content/internal/entity"
)

const notExists = iota

type InfoRepositoryImpl struct {
	db  *redis.Client
	ttl time.Duration
}

func NewInfoRepository(db *redis.Client, ttl time.Duration) *InfoRepositoryImpl {
	return &InfoRepositoryImpl{db, ttl}
}

func (r *InfoRepositoryImpl) Cache(ctx context.Context, key string, value []entity.TitleInfo) error {
	var b bytes.Buffer
	op := "InfoRepository.Cache"

	// Create an encoder and encode the value.
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(value); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Get the bytes from the buffer
	byteSlice := b.Bytes()

	err := r.db.Set(ctx, key, byteSlice, r.ttl).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *InfoRepositoryImpl) Lookup(ctx context.Context, key string) (bool, error) {
	data, err := r.db.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return data != notExists, nil
}

func (r *InfoRepositoryImpl) FromCache(ctx context.Context, key string) ([]entity.TitleInfo, error) {
	op := "InfoRepository.FromCache"

	var ti = make([]entity.TitleInfo, 0)
	var asBytes []byte

	if err := r.db.Get(ctx, key).Scan(&asBytes); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	b := bytes.NewReader(asBytes)

	// Create a decoder and decode the value
	_ = gob.NewDecoder(b).Decode(&ti)

	return ti, nil
}

func (r *InfoRepositoryImpl) Healthcheck(ctx context.Context) bool {
	if err := r.db.Ping(ctx).Err(); err != nil {
		return false
	}

	return true
}
