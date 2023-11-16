package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfig() Config {
	var cfg Config

	if env := os.Getenv("DOCKERIZED"); env == "Yes" {
		data, err := os.ReadFile("config.json")
		if err != nil {
			log.Printf("Failed to read config file: %v\nLoading defaults...", err)
			cfg = defaultConfig()
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			log.Printf("Failed to unmarshal config file into struct: %v\n", err)
			cfg = defaultConfig()
		}
	} else {
		cfg = defaultConfig()
	}

	return cfg
}

// defaults for Docker use
func defaultConfig() Config {
	return Config{Host: "0.0.0.0", Port: "3000"}
}
