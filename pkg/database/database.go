package database

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() error {
	m, err := migrate.New(
		os.Getenv("migrations"),
		os.Getenv("URI"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return m.Up()
}
