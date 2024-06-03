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
