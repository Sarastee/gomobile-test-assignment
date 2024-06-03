package service

import (
	"context"
	"encoding/json"
)

// ExchangeService interface
type ExchangeService interface {
	GetExchangeRateFromAPI(ctx context.Context, val string, date string) (json.RawMessage, error)
	GetCurrenciesFromAPI() (map[string]bool, error)
}

// ExchangeCacheService interface
type ExchangeCacheService interface {
	SetCache(ctx context.Context, val string, data string, content string) error
	GetCache(ctx context.Context, val string, data string) (string, error)
}
