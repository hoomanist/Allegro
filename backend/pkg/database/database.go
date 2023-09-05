package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gopkg.in/ini.v1"
)

func Migrate(SqlCfg *ini.Section) error {
	m, err := migrate.New(
		SqlCfg.Key("migrations").String(),
		SqlCfg.Key("URI").String(),
	)
	if err != nil {
		log.Fatal(err)
	}
	return m.Up()
}
