package cli

import (
	"log"

	"gopkg.in/ini.v1"
)

func Configure(path string) *ini.File {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
