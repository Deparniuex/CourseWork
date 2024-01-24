package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP  ServerConfig `yaml:"http"`
	DB    DBConfig     `yaml:"db"`
	Admin AdminConfig  `yaml:"admin"`
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

type AdminConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Email    string `yaml:"email"`
	Name     string `yaml:"name"`
	Lastname string `yaml:"lastname"`
	Phone    string `yaml:"phone"`
}

func ReadConfig(path string) (*Config, error) {
	cfg := new(Config)
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
