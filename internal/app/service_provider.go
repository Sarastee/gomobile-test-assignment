package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/sarastee/gomobile-test-assignment/internal/api/exchange"
	"github.com/sarastee/gomobile-test-assignment/internal/config"
	"github.com/sarastee/gomobile-test-assignment/internal/config/env"
	"github.com/sarastee/gomobile-test-assignment/internal/service"
	exchangeService "github.com/sarastee/gomobile-test-assignment/internal/service/exchange"
	"github.com/sarastee/platform_common/pkg/closer"
	"github.com/sarastee/platform_common/pkg/db"
	"github.com/sarastee/platform_common/pkg/db/pg"
	"github.com/sarastee/platform_common/pkg/memory_db"
	"github.com/sarastee/platform_common/pkg/memory_db/rs"
)

type serviceProvider struct {
	logger        *zerolog.Logger
	pgConfig      *config.PgConfig
	redisConfig   *config.RedisConfig
	httpConfig    *config.HTTPConfig
	swaggerConfig *config.SwaggerConfig

	dbClient      db.Client
	txManager     db.TxManager
	redisDbClient memory_db.Client

	// currency repo layer

	exchangeService service.ExchangeService

	exchangeImpl *exchange.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Logger ...
func (s *serviceProvider) Logger() *zerolog.Logger {
	if s.logger == nil {
		cfgSearcher := env.NewLogCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Logger config: %s", err.Error())
		}

		s.logger = setupZeroLog(cfg)
	}

	return s.logger
}

func setupZeroLog(logConfig *config.LogConfig) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logConfig.TimeFormat}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger = logger.Level(logConfig.LogLevel)
	zerolog.TimeFieldFormat = logConfig.TimeFormat

	return &logger
}

// PgConfig ...
func (s *serviceProvider) PgConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get PG config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// HTTPConfig ...
func (s *serviceProvider) HTTPConfig() *config.HTTPConfig {
	if s.httpConfig == nil {
		cfgSearcher := env.NewHTTPCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get HTTP config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) RedisConfig() *config.RedisConfig {
	if s.redisConfig == nil {
		cfgSearcher := env.NewRedisCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Redis config:%s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

// SwaggerConfig ...
func (s *serviceProvider) SwaggerConfig() *config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfgSearcher := env.NewSwaggerCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

// DBClient ...
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PgConfig().DSN(), s.Logger())
		if err != nil {
			log.Fatalf("failure while creating DB: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("no connection to DB: %s", err.Error())
		}
		closer.Add(cl.Close)

		log.Printf("DB connected at %s:%d/%s", s.PgConfig().Host, s.PgConfig().Port, s.PgConfig().DbName)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) RedisDBClient(_ context.Context) memory_db.Client {
	if s.redisDbClient == nil {
		redisConfig := s.RedisConfig()
		redisPool := &redis.Pool{
			MaxIdle:     redisConfig.MaxIdle,
			IdleTimeout: redisConfig.IdleTimeout,
			DialContext: func(ctx context.Context) (redis.Conn, error) {
				return redis.DialContext(ctx, "tcp", redisConfig.Address())
			},
			TestOnBorrowContext: func(_ context.Context, conn redis.Conn, lastUsed time.Time) error {
				if time.Since(lastUsed) < time.Minute {
					return nil
				}
				_, err := conn.Do("PING")
				return err
			},
		}
		s.redisDbClient = rs.New(redisPool)

		log.Printf("Redis connected at %s", redisConfig.Address())

		closer.Add(s.redisDbClient.Close)
	}

	return s.redisDbClient
}

// TxManager ...
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = pg.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ExchangeService() service.ExchangeService {
	if s.exchangeService == nil {
		s.exchangeService = exchangeService.NewService(
			s.Logger(),
		)
	}

	return s.exchangeService
}

func (s *serviceProvider) ExchangeImpl() *exchange.Implementation {
	if s.exchangeImpl == nil {
		s.exchangeImpl = exchange.NewImplementation(
			s.Logger(),
			s.ExchangeService())
	}

	return s.exchangeImpl
}
