package models

import (
	"time"

	"github.com/hiteshnayak305/go-rest-project/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"

	rows, err := db.SqlConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event = make([]Event, 0)
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (e *Event) CreateEvent() error {

	query := `
	INSERT INTO events(name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.SqlConnection.Prepare(query)
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
	return err
}
