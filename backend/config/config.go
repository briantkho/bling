package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Address     string `yaml:"address"`
		Environment string `yaml:"environment"`
	} `yaml:"server"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
