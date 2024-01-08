package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP ServerConfig `yaml:"http"`
	DB   DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port    string        `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password" env:"PASSWORD"`
	DBName   string `yaml:"dbname"`
}

func ReadConfig(path string) (*Config, error) {
	cfg := new(Config)
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
