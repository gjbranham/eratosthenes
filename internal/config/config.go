package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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
			gin.DefaultWriter.Write([]byte(fmt.Sprintf("Failed to read config file: %v\nLoading defaults...", err)))

			cfg = defaultConfig()
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			gin.DefaultWriter.Write([]byte(fmt.Sprintf("Failed to unmarshal config file into struct: %v\n", err)))

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
