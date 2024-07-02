package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

type Fact struct {
	id       int
	question string
	answer   string
	sentence string
	qType    string
}

type Source struct {
	id     int
	author string
	link   string
	book   string
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type DB struct {
	db   *sql.DB
	rows int
}

var db DB

func init() {
	var err error
	fmt.Printf("Hi")

	db.db, err = sql.Open("sqlite3", "facts.db")
	if err != nil {
		log.Println(err)
	}

	rows, err := db.db.Query("SELECT id FROM Facts ORDER BY id DESC LIMIT 1")
	var tempFact Fact
	for rows.Next() {
		rows.Scan(&tempFact.id)
	}
	db.rows = tempFact.id
}

func GetRandQuestion() (Fact, error) {
	// TODO: turn this into a random poll that runs one with a random 10 item selection, one that selects one of the random items
	v := rand.Intn(db.rows)
	preState, err := db.db.Prepare("SELECT id, question, answer FROM facts where id = ?")

	if err != nil {
		return Fact{}, err
	}

	rows, err := preState.Query(v)
	var tempFact Fact
	for rows.Next() {
		rows.Scan(&tempFact.id, &tempFact.question, &tempFact.answer)
		return tempFact, nil
	}
	return Fact{}, errors.New("something went wrong while selecting question")
}

func GetSource(objID int) (Source, error) {
	// THIS ISNT GONNA WoRK, NEED TO USE THE FORIGEN ID NOT THE ID
	preState, err := db.db.Prepare("SELECT * FROM sources where id = ?")
	if err != nil {
		return Source{}, err
	}

	rows, err := preState.Query(objID)
	var tempSoure Source
	for rows.Next() {
		rows.Scan(&tempSoure.id, &tempSoure.author, &tempSoure.book, &tempSoure.link)
		return tempSoure, nil
	}
	return Source{}, errors.New("something went wrong while selecting source")
}

// func setSource({obj}){

// }

// func addQuestion({obj}){

// }
