package configs

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)


type Configs struct {
	App              App                    `envPrefix:"APP_"`
	Http             Http                   `envPrefix:"HTTP_"`
}


func Configure() (*Configs, error) {
	err := godotenv.Load(".env")
	if err != nil {
		slog.With("err", err.Error()).Error("reading .env file error")
	}

	cfg := &Configs{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing configuration error: %w", err)
	}

	return cfg, nil
}