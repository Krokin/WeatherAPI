package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool  `yaml:"is_debug"`
	OWN_API string `yaml:"own_api"`
	Port string `yaml:"port"`

}

func GetConfig() (*Config, error) {
	var cfg = new(Config)
	err := cleanenv.ReadConfig("config.yml", &cfg)
	return cfg, err
}