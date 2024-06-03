package exchange

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/gomobile-test-assignment/internal/service"
)

const (
	DateParam = "date" // DateParam is constant for date query param
	ValParam  = "val"  // ValParam is constant for val query param
)

// Implementation struct
type Implementation struct {
	logger               *zerolog.Logger
	exchangeService      service.ExchangeService
	exchangeCacheService service.ExchangeCacheService
}

// NewImplementation function which creates Implementation object
func NewImplementation(logger *zerolog.Logger, exchangeService service.ExchangeService, exchangeCacheService service.ExchangeCacheService) *Implementation {
	return &Implementation{
		logger:               logger,
		exchangeService:      exchangeService,
		exchangeCacheService: exchangeCacheService,
	}
}
