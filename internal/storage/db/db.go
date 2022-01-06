package db

import (
	"calendar/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Database struct {
	DB *sqlx.DB
}

type EventDB struct {
	Id string `json:"id"`
	models.Event
}

func NewDb(dsn string) (*Database, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{
		DB: db,
	}, nil
}

func (db *Database) CreateEvent(title, date string) (EventDB, error) {
	var event EventDB

	err := db.DB.QueryRow("insert into accounts (title, date) values ($1,$2) returning id,title,date", title, date).Scan(&event.Id, &event.Title, &event.Date)
	if err != nil {
		return EventDB{}, err
	}
	return event, nil
}

func (db *Database) UpdateEvent(id, title, date string) (EventDB, error) {
	var event EventDB

	err := db.DB.QueryRow("update accounts set title = $1, date = $2 where id = $3 returning id,title,date", title, date, id).Scan(&event.Id, &event.Title, &event.Date)
	if err != nil {
		return EventDB{}, err
	}

	return event, nil
}

func (db *Database) DeleteEvent(id string) error {
	_, err := db.DB.Exec("delete from accounts where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) EventForDay() ([]EventDB, error) {
	var event []EventDB
	today := time.Now()
	todayFormatted := today.String()[0:10]

	result, err := db.DB.Query("SELECT * FROM accounts WHERE date::date = $1", todayFormatted)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		var ev EventDB
		err := result.Scan(&ev.Id, &ev.Title, &ev.Date)
		if err != nil {
			return nil, err
		}
		event = append(event, ev)
	}
	log.Println(event)
	return event, nil
}
