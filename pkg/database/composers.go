package database

import (
	"database/sql"

	"gopkg.in/ini.v1"
)

type Composer struct {
	id            int    `json:"id"`
	composer_name string `json:"name"`
	description   string `json:"description"`
	birth         int    `json:"birth"`
	death         int    `json:"death"`
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
		var id int
		var name string
		var description string
		var birth int
		var death int
		err = rows.Scan(&id, &name, &description, &birth, &death)
		if err != nil {
			return nil, err
		}
		composers = append(composers, Composer{
			id:            id,
			composer_name: name,
			description:   description,
			birth:         birth,
			death:         death,
		})
	}
	return composers, nil
}
