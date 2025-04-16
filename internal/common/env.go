package common

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func InitEnv() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(".env."+env, ".env")
	if err != nil {
		log.Fatal().Err(err)
	}
}

func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
