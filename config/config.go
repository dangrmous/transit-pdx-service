package config

import "os"

type Config struct {
	// Add configuration fields here, for example:
	Port int
	Host string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetTrimetApiKey() string {
	return os.Getenv("TRIMET_API_KEY")
}
