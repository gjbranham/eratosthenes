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
		fo, err := os.Open("config.json")
		if err != nil {
			log.Printf("Failed to open config file: %v\nLoading defaults...", err)
			cfg = Config{Host: "0.0.0.0", Port: "3000"} // Our Dockerized defaults
		}
		defer fo.Close()
		jsonParser := json.NewDecoder(fo)
		jsonParser.Decode(&cfg)
	} else {
		cfg = Config{Host: "0.0.0.0", Port: "3000"}
	}

	return cfg
}
