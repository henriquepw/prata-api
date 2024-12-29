package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/henriquepw/pobrin-api/pkg/validate"
	"github.com/joho/godotenv"
)

type Config struct {
	TZ           string `validate:"required,uppercase"`
	Port         string `validate:"required,numeric"`
	JWTSecret    string `validate:"required"`
	DatabaseURL  string `validate:"required"`
	DatabaseName string `validate:"required"`
}

var config *Config

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "../..")

	env := os.Getenv("ENV")
	if env == "" {
		env = ".env"
	}

	godotenv.Load(filepath.Join(rootDir, env))
	Load()
}

// Load carrega as configurações da variável de ambiente
func Load() {
	config = &Config{
		TZ:           os.Getenv("TZ"),
		Port:         os.Getenv("PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}

	if config.Port == "" {
		config.Port = "3333"
	}

	err := validate.Check(config)
	if err != nil {
		log.Panic(err)
	}
}

// Env retorna a configuração
func Env() *Config {
	if config == nil {
		Load()
	}

	return config
}
