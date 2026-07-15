package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const configPath = ".env"

type Configuration struct {
	HTTPServer HTTPServerConfig
	//Gateway GatewayConfig
}

type HTTPServerConfig struct {
	Address             string   `env:"HTTP_SERVER_ADDRESS"`
	ReadTimeoutSeconds  uint     `env:"HTTP_SERVER_READ_TIMEOUT" envDefault:"30"`
	WriteTimeoutSeconds uint     `env:"HTTP_SERVER_WRITE_TIMEOUT" envDefault:"30"`
	CORSAllowedOrigins  []string `env:"HTTP_SERVER_CORS_ORIGINS" envDefault:"http://localhost:3000"`
}

func New() (*Configuration, error) {
	return LoadAndParseConfig(configPath)
}

func LoadAndParseConfig(path string) (*Configuration, error) {
	if err := loadEnvFile(path); err != nil {
		return nil, err
	}

	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse env config: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

func loadEnvFile(path string) error {
	if err := godotenv.Load(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("load env file %s: %w", path, err)
	}
	return nil
}

func (c *Configuration) validate() error {
	var errs []error

	if c.HTTPServer.Address == "" {
		errs = append(errs, errors.New("token: HTTP_SERVER_ADDRESS is required"))
	}

	return errors.Join(errs...)
}
