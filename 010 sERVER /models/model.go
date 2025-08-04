package models

import (
	"time"

	"http.com/com/db"
)

type Event struct {
	ID          int64  `binding:"required"`
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(name,description,location,datetime,user_id) 
	VALUE (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err

	}
	id, err := result.LastInsertId()
	e.ID = id

	// events = append(events, e)
	return err
}

func GetAllEvents() []Event {
	query := `SELECT * FROM events`
	db.DB.Query(query)

	return events
}
