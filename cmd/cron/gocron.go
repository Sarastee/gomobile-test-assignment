package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/sarastee/gomobile-test-assignment/internal/app"
	"github.com/sarastee/gomobile-test-assignment/internal/config"
	"github.com/sarastee/gomobile-test-assignment/internal/config/env"
	"github.com/sarastee/gomobile-test-assignment/internal/repository"
	"github.com/sarastee/gomobile-test-assignment/internal/repository/exchange"
	"github.com/sarastee/platform_common/pkg/closer"
	"github.com/sarastee/platform_common/pkg/db/pg"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

func main() { // nolint: revive // this is a small executable file and big length of function is possible
	ctx := context.Background()
	// Config
	err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Logger config
	LogCfgSearcher := env.NewLogCfgSearcher()
	logCfg, err := LogCfgSearcher.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Logger
	l := app.SetupZeroLog(logCfg)

	err = config.Load(configPath)
	if err != nil {
		l.Fatal().Err(err)
	}

	// Postgres config
	PgCfgSearcher := env.NewPgCfgSearcher()
	pgCfg, err := PgCfgSearcher.Get()
	if err != nil {
		l.Fatal().Msgf("unable to get PG config: %s", err.Error())
	}

	// Postgres
	cl, err := pg.New(ctx, pgCfg.DSN(), l)
	if err != nil {
		log.Fatalf("failure while creating DB: %v", err)
	}

	err = cl.DB().Ping(ctx)
	if err != nil {
		log.Fatalf("no connection to DB: %s", err.Error())
	}
	closer.Add(cl.Close)

	log.Printf("DB connected at %s:%d/%s", pgCfg.Host, pgCfg.Port, pgCfg.DbName)

	// Repository
	exchangeRepository := exchange.NewExchangeRepo(l, cl)

	// Cron
	cronScheduler, err := gocron.NewScheduler(gocron.WithLocation(time.UTC))

	if err != nil {
		l.Fatal().Msgf("INIT: start cronScheduler error: %s", err)
	}

	if _, err = cronScheduler.NewJob(
		gocron.WeeklyJob(0,
			gocron.NewWeekdays(time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday),
			gocron.NewAtTimes(
				gocron.NewAtTime(16, 14, 0)),
		),
		gocron.NewTask(
			func(rep repository.ExchangeRepository, l *zerolog.Logger) {
				if err := rep.InsertDailyData(ctx); err != nil {
					l.Error().Msgf("ERROR: %s", errors.Wrap(err, "in cron job of inserting daily data"))
				} else {
					l.Info().Msgf("daily data inserted successfully")
				}
			},
			exchangeRepository,
			l,
		),
	); err != nil {
		l.Fatal()
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	cronScheduler.Start()
	l.Println("Start: service started")

	s := <-interrupt
	l.Printf("RUN - signal: %s", s.String())

	// Shutdown
	err = cronScheduler.Shutdown()
	if err != nil {
		l.Fatal().Msgf("STOP - cronScheduler.Shutdown: %s", err)
	}

	l.Println("Stop - service stopped")
}
