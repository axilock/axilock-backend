package util

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBSource            string        `env:"DB_SOURCE"`
	HTTPServerAddress   string        `env:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress   string        `env:"GRPC_SERVER_ADDRESS"`
	AccessTokenDuration time.Duration `env:"ACCESS_TIME_DURATION" envDefault:"1h"`
	RunningEnv          string        `env:"RUNNING_ENV" envDefault:"dev"`
	RedisAddr           string        `env:"REDIS_ADDR"`
	RedisPassword       string        `env:"REDIS_PASS"`
	GithubAppID         string        `env:"GITHUB_APP_ID"`
	GithubAppPrivateKey string        `env:"GITHUB_PRIV_KEY,file" envDefault:"axilock.pem"`
	GithubClientSecret  string        `env:"GITHUB_CLIENT_SECRET"`
	GithubClientID      string        `env:"GITHUB_CLIENT_ID"`
	DiscordWebhook      string        `env:"DISCORD_WEBHOOK" envDefault:""`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = env.Parse(&config)
	return
}
