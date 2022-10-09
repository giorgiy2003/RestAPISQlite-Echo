package repository

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Connection *sql.DB

func OpenTable() error {
	var err error
	if _, err = os.Stat("SQliteDatabase.db"); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create("SQliteDatabase.db")
		if err != nil {
			return err
		}
	}
	Connection, err = sql.Open("sqlite3", "./SQliteDatabase.db")
	if err != nil {
		return err
	}
	table, err := Connection.Exec(`CREATE TABLE IF NOT EXISTS person
	(
		person_id INTEGER PRIMARY KEY AUTOINCREMENT, 
		person_email TEXT NOT NULL,
		person_phone TEXT NOT NULL,
		person_firstName TEXT NOT NULL,
		person_lastName TEXT NOT NULL
	)`)
	if err != nil {
		return err
	}
	log.Println(table.RowsAffected())
	return nil
}
