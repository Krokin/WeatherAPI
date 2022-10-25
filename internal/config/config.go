package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	OWN_API string `yaml:"OWN_API"`
	Port string `yaml:"port"`
}

func GetConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("../../configs/config.yml", &cfg)
	return &cfg, err
}