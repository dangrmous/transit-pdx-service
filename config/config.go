package config

import "os"

type Config struct {
	// Add configuration fields here, for example:
	Port         int
	Host         string
	TrimetApiKey string
}

func NewConfig() *Config {
	return &Config{
		Port:         8000,
		Host:         "localhost",
		TrimetApiKey: os.Getenv("TRIMET_API_KEY"),
	}
}
