package main

import (
	"context"
	"flag"
	"log"

	"github.com/sarastee/gomobile-test-assignment/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

// @title Currency Service
// @version 1.0.0
// @description Сервис для получения курса валют
// @schemes http

// @contact.name Ilya Lyakhov
// @contact.email ilja.sarasti@mail.ru

// @host localhost:8082
// @BasePath /

func main() {
	ctx := context.Background()

	application, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("init app failure: %s", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("failure while running the application: %s", err)
	}
}
