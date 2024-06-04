package exchange

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/gomobile-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	currenciesTable = "currencies"

	valuteIDColumn      = "valute_id"
	dateColumn          = "date"
	numericCodeColumn   = "numeric_code"
	characterCodeColumn = "character_code"
	nominalColumn       = "nominal"
	nameColumn          = "name"
	valueColumn         = "value"
	vunitRateColumn     = "vunit_rate"
)

var (
	BaseURL            = "https://www.cbr.ru/scripts/"         // BaseURL variable
	FindByDateEndpoint = BaseURL + "XML_daily.asp?date_req=%v" // FindByDateEndpoint variable
)

var _ repository.ExchangeRepository = (*Repo)(nil)

// Repo ...
type Repo struct {
	logger *zerolog.Logger
	db     db.Client
}

// NewExchangeRepo get new repo instance
func NewExchangeRepo(logger *zerolog.Logger, dbClient db.Client) *Repo {
	return &Repo{
		logger: logger,
		db:     dbClient,
	}
}
