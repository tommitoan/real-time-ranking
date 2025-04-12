package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var Config Service

type (
	Service struct {
		App   App   `yaml:"app"`
		Redis Redis `yaml:"redis"`
	}
	App struct {
		LoggingLevel string `yaml:"logging_level"`
		Port         string `yaml:"port"`
	}
	Redis struct {
		ConnectionString string `yaml:"connection_string"`
	}
)

func Load(path string) (*Service, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Service
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	Config = cfg
	return &cfg, nil
}
