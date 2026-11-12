package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AiDEX     AiDEXConfig
	Nightscout NightscoutConfig
	PollInterval time.Duration
}

type AiDEXConfig struct {
	APIURL  string
	UserID  string
	Token   string
	Brand   string
	AppInfo string
	UserAgent string
}

type NightscoutConfig struct {
	URL       string
	APISecret string
}

func Load() Config {
	godotenv.Load()
	pollInterval := 1 * time.Minute
	if v := os.Getenv("POLL_INTERVAL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			pollInterval = d
		}
	}

	return Config{
		AiDEX: AiDEXConfig{
			APIURL:    getEnv("AIDEX_API_URL", "https://china.pancares.com/backend/aidex-x/cgmRecord/getCgmRecordsByPageInfo"),
			UserID:    getEnv("AIDEX_USER_ID", ""),
			Token:     getEnv("AIDEX_TOKEN", ""),
			Brand:     getEnv("AIDEX_BRAND", "AiDEX-X"),
			AppInfo:   getEnv("AIDEX_APP_INFO", "com.microtechmd.AiDEX-X"),
			UserAgent: getEnv("AIDEX_USER_AGENT", "EU_mmol/2.4.1 (com.microtechmd.AiDEX-X; build:26032001; iOS 26.4.0) Alamofire/5.6.4"),
		},
		Nightscout: NightscoutConfig{
			URL:       getEnv("NIGHTSCOUT_URL", "http://localhost:1337"),
			APISecret: getEnv("NIGHTSCOUT_API_SECRET", ""),
		},
		PollInterval: pollInterval,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
