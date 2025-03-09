package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database struct {
		User    string `yaml:"user"`
		DBName  string `yaml:"dbname"`
		SSLMode string `yaml:"sslmode"`
	} `yaml:"database"`
	Server struct {
		Protocol string `yaml:"protocol"`
		Port     string `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("ошибка чтения файла %v", path)
	}

	if err = yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalf("ошибка преобразования yaml: %v", err)
	}
}
