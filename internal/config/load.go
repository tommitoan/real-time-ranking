package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var Config Service

type (
	Service struct {
		App      App      `yaml:"app"`
		Redis    Redis    `yaml:"redis"`
		Database Database `yaml:"database"`
	}
	App struct {
		LoggingLevel string `yaml:"logging_level"`
		Port         string `yaml:"port"`
	}
	Redis struct {
		ConnectionString string `yaml:"connection_string"`
	}
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		DSN      string `yaml:"dsn"`
		Debug    bool   `yaml:"debug"`
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
