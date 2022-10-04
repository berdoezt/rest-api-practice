package database

import (
	"fmt"
	"rest-api-practice/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	var host = "localhost"
	var port = 5432
	var username = "postgres"
	var password = "password"
	var dbName = "db-go-sql"

	var conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(model.Person{})
	if err != nil {
		fmt.Println("error on migration", err)
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}

func (d Database) GetPersons() ([]model.Person, error) {
	dbg := d.db.Find(&[]model.Person{})

	rows, err := dbg.Rows()
	if err != nil {
		return nil, err
	}

	persons := make([]model.Person, 0)

	for rows.Next() {
		var person model.Person

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			continue
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (d Database) CreatePerson(person model.Person) (model.Person, error) {
	dbg := d.db.Create(&person)

	row := dbg.Row()

	var personResult model.Person

	err := row.Scan(&personResult.ID, &personResult.FirstName, &personResult.LastName)

	return personResult, err
}
