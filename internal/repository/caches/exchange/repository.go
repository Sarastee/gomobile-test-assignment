package exchange

import (
	"github.com/sarastee/gomobile-test-assignment/internal/config"
	"github.com/sarastee/gomobile-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/memory_db"
)

const (
	setCommand = "SET"
	getCommand = "GET"
	exCommand  = "EX"
)

var _ repository.ExchangeCacheRepository = (*ExchangeCacheRepo)(nil)

type ExchangeCacheRepo struct {
	client      memory_db.Client
	redisConfig *config.RedisConfig
}

func NewExchangeCacheRepo(client memory_db.Client, redisConfig *config.RedisConfig) *ExchangeCacheRepo {
	return &ExchangeCacheRepo{
		client:      client,
		redisConfig: redisConfig,
	}
}
