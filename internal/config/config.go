package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	PostgresConfig
	RedisConfig
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST",required`
	Port     string `env:"POSTGRES_PORT",required`
	User     string `env:"POSTGRES_USER",required`
	Password string `env:"POSTGRES_PASSWORD",required`
	DBName   string `env:"POSTGRES_DB_NAME",required`
}

type RedisConfig struct {
}

func MustLoad() *Config {
	var postgresConfig PostgresConfig
	var redisConfig RedisConfig

	err := env.Parse(&postgresConfig)
	if err != nil {
		panic(err)
	}

	err = env.Parse(&redisConfig)
	if err != nil {
		panic(err)
	}

	return &Config{
		PostgresConfig: postgresConfig,
		RedisConfig:    redisConfig,
	}
}
