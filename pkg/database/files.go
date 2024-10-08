package database

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type File struct {
	Uploadtime time.Time
	Filepath   string
}

func NewFile(file *File) error {
	db, err := sql.Open("postgres", os.Getenv("URI"))
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
