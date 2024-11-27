package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	GMPKey string `json:"gmp_key"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open("./etc/config.json")
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("failed to decode config file: %w", err)
	}

	return config, nil
}
