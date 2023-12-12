package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env  string `yaml:"env" env-default:"local"`
	Port string `yaml:"port" env-defaul:"8000"`
	Postgres
}

type Postgres struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	DB_port  string `env:"DB_PORT" env-default:"5432"`
	Username string `env:"POSTGRES_USER" env-default:"root"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	DBName   string `env:"DB_NAME" env-default:"chat"`
	SSLMode  string `env:"SSL_MODE" env-default:"disable"`
}

func MustLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	configPath := os.Getenv("CONFIG_PATH") // Другой способ: можно брать из флага при запуске приложения
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}
