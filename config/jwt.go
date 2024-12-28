package config

import (
	"os"
	"time"
)

type JWTConfig struct {
	Secret               []byte
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

var JWT *JWTConfig

func InitJWTConfig() {
	JWT = &JWTConfig{
		Secret:               []byte(getEnvOrDefault("JWT_SECRET", "jfg9w8394roeqf298yr9rewjflwjj109303")),
		AccessTokenDuration:  time.Hour * 24,     // 1 day
		RefreshTokenDuration: time.Hour * 24 * 7, // 7 days
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
