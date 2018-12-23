package config

import (
	"fmt"
	"os"
)

const mongoAddrEnvVariable = "MONGO_URL"

type Config struct {
	mongoURL string
}

type Contract interface {
	MongoURL() string
}

func InitConfig() (*Config, error) {

	config := &Config{}

	mongoURL := os.Getenv(mongoAddrEnvVariable)

	if mongoURL == "" {
		return nil, fmt.Errorf("env variable [%s] is empty", mongoAddrEnvVariable)
	}

	config.mongoURL = mongoURL

	return config, nil
}

func (c *Config) MongoURL() string {
	return c.mongoURL
}
