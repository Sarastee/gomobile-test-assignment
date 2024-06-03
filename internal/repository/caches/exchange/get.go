package exchange

import (
	"context"
	"errors"

	"github.com/gomodule/redigo/redis"
	"github.com/sarastee/gomobile-test-assignment/internal/repository"
)

func (e *ExchangeCacheRepo) GetCache(ctx context.Context, key string) (string, error) {
	db := e.client.DB()
	content, err := db.String(db.DoContext(ctx, getCommand, key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return "", repository.ErrCacheNotFound
		}

		return "", err
	}
	return content, nil
}
