package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
	Password     string    `json:"password"`
}

func NewUser(user User) error {
	db, err := sql.Open("postgres", os.Getenv("URI"))
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = db.Exec("INSERT INTO users (name, creationdate, password) VALUES ($1, $2, $3)",
		user.Name,
		user.CreationDate,
		user.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUser(name string) (string, error) {
	db, err := sql.Open("postgres", os.Getenv("URI"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	var password string
	row := db.QueryRow("SELECT password FROM users WHERE name=$1", name)
	err = row.Scan(&password)
	if err != nil {
		fmt.Println("Hoo")
		return "", err
	}
	return password, nil
}

// only exists for testing purposes
func GetUsers() ([]User, error) {
	db, err := sql.Open("postgres", os.Getenv("URI"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var (
			id            int
			name          string
			creation_date time.Time
			password      string
		)
		err = rows.Scan(&id, &name, &creation_date, &password)
		if err != nil {
			return nil, err
		}
		users = append(users, User{
			Id:           id,
			Name:         name,
			CreationDate: creation_date,
			Password:     password,
		})
	}
	return users, nil
}
