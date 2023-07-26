package cli

import (
	"log"

	"gopkg.in/ini.v1"
)

type Configuration struct {
	Port string
	DB   *ini.Section
}

func Configure(path string) *Configuration {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return &Configuration{
		Port: cfg.Section("").Key("port").String(),
		DB:   cfg.Section("database"),
	}
}
