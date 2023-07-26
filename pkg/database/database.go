package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"gopkg.in/ini.v1"
)

func Migrate(SqlCfg *ini.Section) error {
	db, err := sql.Open("postgres", SqlCfg.Key("URI").String())
	if err != nil {
		log.Fatal("can't connect to database")
	}
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		SqlCfg.Key("migrations").String(),
		"postgres", driver,
	)
	if err != nil {
		return err
	}
	return m.Up()
}
