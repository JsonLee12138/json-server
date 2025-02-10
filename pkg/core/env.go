package core

import (
	"os"
	"strings"
	"sync"
)

const ENV_MODE_KEY = "JSON_ENV_MODE"

type ENV_MODE string

const (
	DevMode  ENV_MODE = "development"
	ProMode  ENV_MODE = "production"
	TestMode ENV_MODE = "test"
)

var (
	currentEnv ENV_MODE
	modeOnce   sync.Once
)

func ParseEnv(env string) ENV_MODE {
	normalizedEnv := strings.ToLower(strings.TrimSpace(env))
	switch normalizedEnv {
	case "development", "dev", "":
		return DevMode
	case "production", "prod", "pro":
		return ProMode
	case "test", "testing":
		return TestMode
	default:
		return DevMode
	}
}

func Mode() ENV_MODE {
	if currentEnv == "" {
		modeOnce.Do(func() {
			currentEnv = ParseEnv(os.Getenv(ENV_MODE_KEY))
			if currentEnv == "" {
				currentEnv = DevMode
			}
		})
	}
	return currentEnv
}

func SetMode(mode ENV_MODE) {
	os.Setenv(ENV_MODE_KEY, string(mode))
}
