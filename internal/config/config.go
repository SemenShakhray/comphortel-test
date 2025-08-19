package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const configPath = ".env"

type Config struct {
	MigrationDir string `env:"MIGRATION_DIR" env-default:"./migrations"`

	DB struct {
		Host     string `env:"DB_HOST"     env-default:"postgres"`
		Port     int    `env:"DB_PORT"     env-default:"5432"`
		User     string `env:"DB_USER"     env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
		Name     string `env:"DB_NAME"     env-default:"documents"`
	}

	Server struct {
		Host        string        `env:"SERVER_HOST"          env-default:"0.0.0.0"`
		Port        int           `env:"SERVER_PORT"          env-default:"8081"`
		Timeout     time.Duration `env:"SERVER_TIMEOUT"       env-default:"5s"`
		IdleTimeout time.Duration `env:"SERVER_IDDLE_TIMEOUT" env-default:"60s"`
	}
}

func MustLoad() *Config {
	if err := godotenv.Load(configPath); err != nil {
		log.Fatalf("cannot load env file: %s", err)
	}

	log.Println("using env file: " + configPath)

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("failed to read env: " + err.Error())
	}

	return &cfg
}
