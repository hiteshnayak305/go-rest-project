package models

import (
	"time"

	"github.com/hiteshnayak305/go-rest-project/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

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

func GetEventByID(id int64) (*Event, error) {
	var event Event
	query := "SELECT * FROM events WHERE id = ?"
	row := db.SqlConnection.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent() error {

	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	`
	stmt, err := db.SqlConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (event Event) DeleteEvent() error {

	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.SqlConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) Register(userId int64) error {

	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := db.SqlConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.ID, userId)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}

func (e Event) CancelRegistration(userId int64) error {

	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.SqlConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}
