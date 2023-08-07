package database

import (
	"database/sql"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type File struct {
	Uploadtime time.Time
	Filepath   string
}

func NewFile(SqlCfg *ini.Section, file *File) error {
	db, err := sql.Open("postgres", SqlCfg.Key("URI").String())
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = db.Exec("INSERT INTO files VALUES ($1, $2)", file.Uploadtime, file.Filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
