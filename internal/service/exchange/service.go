package exchange

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/gomobile-test-assignment/internal/repository"
	"github.com/sarastee/gomobile-test-assignment/internal/service"
)

var (
	BaseURL               = "https://www.cbr.ru/scripts/"         // BaseURL variable
	FindByDateEndpoint    = BaseURL + "XML_daily.asp?date_req=%v" // FindByDateEndpoint variable
	CurrencyCodesEndpoint = BaseURL + "XML_valFull.asp"           // CurrencyCodesEndpoint variable
)

var _ service.ExchangeService = (*Service)(nil)

// Service exchange struct
type Service struct {
	logger       *zerolog.Logger
	exchangeRepo repository.ExchangeRepository
}

// NewExchangeService creates new service struct
func NewExchangeService(logger *zerolog.Logger, exchangeRepository repository.ExchangeRepository) *Service {
	return &Service{
		logger:       logger,
		exchangeRepo: exchangeRepository,
	}
}
