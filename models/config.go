package models

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	AdoApi           string `json:"adoApi"`
	AdoToken         string `json:"adoToken"`
	ConnectionString string `json:"connectionString"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config
	configFile, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	byteValue, err := io.ReadAll(configFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteValue, &config)
	return config, err
}
