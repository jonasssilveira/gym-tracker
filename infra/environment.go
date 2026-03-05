package config

import (
	"flag"
)

func Parse() Config {
	var environment string
	flag.StringVar(&environment, "env", "local", "default environment")
	flag.Parse()
	cfg := Config{}
	if environment == "local" {
		cfg = CreateConfig("infra/application.yaml")
	}
	return cfg
}
