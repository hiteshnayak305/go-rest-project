package models

import "github.com/hiteshnayak305/go-rest-project/db"

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) CreateUser() error {

	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)`

	stmt, err := db.SqlConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}
