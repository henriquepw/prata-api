package config

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/henriquepw/pobrin-api/pkg/validate"
	"github.com/joho/godotenv"
)

type Config struct {
	TZ          string `validate:"required,uppercase"`
	Port        string `validate:"required,numeric"`
	JWTSecret   string `validate:"required"`
	DatabaseURL string `validate:"required"`
}

var (
	config      *Config
	configMutex = &sync.Mutex{}
)

func init() {
	if os.Getenv("ENV") == "test" {
		return
	}

	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "../..")

	env := os.Getenv("ENV")
	if env == "" {
		env = ".env"
	}

	godotenv.Load(filepath.Join(rootDir, env))
	load()
}

func load() {
	config = &Config{
		TZ:          os.Getenv("TZ"),
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if config.Port == "" {
		config.Port = "3333"
	}

	err := validate.Check(config)
	if err != nil {
		panic(err)
	}
}

func Env() *Config {
	if config == nil {
		configMutex.Lock()
		defer configMutex.Unlock()

		if config == nil {
			load()
		}
	}

	return config
}
