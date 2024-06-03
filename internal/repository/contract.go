package repository

import "context"

// ExchangeCacheRepository interface
type ExchangeCacheRepository interface {
	SetCache(ctx context.Context, key string, content string) error
	GetCache(ctx context.Context, key string) (string, error)
}
