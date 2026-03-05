package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   `yaml:"server"`
	Database `yaml:"database"`
	Telegram `yaml:"telegram"`
}

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Database struct {
	Cache      `yaml:"cache"`
	Postgresql `yaml:"postgresql"`
}

type Cache struct {
	Redis `yaml:"redis"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DbName   int    `yaml:"db_name"`
}

type Postgresql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Telegram struct {
	Token      string `yaml:"token"`
	Debug      bool   `yaml:"debug"`
	Timeout    int    `yaml:"timeout"`
	NewUpdated int    `yaml:"new_updated"`
}

func CreateConfig(applicationFilename string) Config {
	yamlFile, err := os.ReadFile(applicationFilename)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	// Access the configuration values
	fmt.Printf("Server Port: %d\n", cfg.Server.Port)
	return cfg
}
