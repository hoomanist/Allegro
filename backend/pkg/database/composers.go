package database

import (
	"database/sql"
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type Composer struct {
	Id            int    `json:"id"`
	Composer_name string `json:"name"`
    Photo string `json:"photo"`
	Description   string `json:"desc"`
	Birth         int    `json:"birth"`
	Death         int    `json:"death"`
}

func ListComposers(SqlCfg *ini.Section) ([]Composer, error) {
	db, err := sql.Open("postgres", SqlCfg.Key("URI").String())
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM composers")
	if err != nil {
		return nil, err
	}
	composers := []Composer{}
	for rows.Next() {
		var (
			id    int
			birth int
			death int

            photo string
			name        string
			description string
		)
		err = rows.Scan(&id, &name, &photo, &description, &birth, &death)
		if err != nil {
			return nil, err
		}
		composers = append(composers, Composer{
			Id:            id,
			Composer_name: name,
            Photo:         photo,
			Description:   description,
			Birth:         birth,
			Death:         death,
		})
	}
	return composers, nil
}

func NewComposer(SqlCfg *ini.Section, composer *Composer) error {
	db, err := sql.Open("postgres", SqlCfg.Key("URI").String())
	if err != nil {
		log.Println("hallo")
		return err
	}
	cmd := fmt.Sprintf("INSERT INTO composers (composer_name, photo, description, birth, death) VALUES ('%s','%s', '%s', '%d', '%d');",
		composer.Composer_name, composer.Photo, composer.Description, composer.Birth, composer.Death)

	_, err = db.Exec(cmd)
	if err != nil {
		log.Println("Bonjour")
		return err
	}
	return nil
}
