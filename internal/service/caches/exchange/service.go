package exchange

import (
	"github.com/sarastee/gomobile-test-assignment/internal/repository"
	"github.com/sarastee/gomobile-test-assignment/internal/service"
)

var _ service.ExchangeCacheService = (*Service)(nil)

type Service struct {
	exchangeCacheRepo repository.ExchangeCacheRepository
}

func NewService(exchangeCacheRepository repository.ExchangeCacheRepository) *Service {
	return &Service{
		exchangeCacheRepo: exchangeCacheRepository,
	}
}
