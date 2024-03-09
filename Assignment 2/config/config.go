package config

import (
	"fmt"
	"hacktiv-assignment-2/utils/common"
	"os"
)

type DBConfig struct {
	Url string
}

type APIConfig struct {
	APIHost, APIPort string
}

type Config struct {
	APIConfig
	DBConfig
}

func (c *Config) ReadConfig() error {
	err := common.LoadENV()
	if err != nil {
		return err
	}

	c.DBConfig = DBConfig{
		Url: os.Getenv("DB_URL"),
	}

	c.APIConfig = APIConfig{
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}

	if c.DBConfig.Url == "" || c.APIConfig.APIHost == "" || c.APIConfig.APIPort == "" {
		return fmt.Errorf("missing required enivronment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
