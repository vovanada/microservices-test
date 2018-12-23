package config

import (
	"fmt"
	"os"
)

const portServiceAddrEnvVariable = "SERVICE_PORT_ADDR"

type Config struct {
	portServiceAddr string
}

type Contract interface {
	PortServiceAddr() string
}

func InitConfig() (*Config, error) {

	config := &Config{}

	portServiceAddr := os.Getenv(portServiceAddrEnvVariable)

	if portServiceAddr == "" {
		return nil, fmt.Errorf("env variable [%s] is empty", portServiceAddrEnvVariable)
	}

	config.portServiceAddr = portServiceAddr

	return config, nil
}

func (c *Config) PortServiceAddr() string {
	return c.portServiceAddr
}
