package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var db, err = sql.Open("sqlite3", "data/database/poll.db")
	if err != nil {
		log.Fatalln("Error opening database: ", err)
	}
	DB = db

	createQuestionTable()
	createAnswerTable()
}
