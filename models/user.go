package models

import (
	"github.com/hiteshnayak305/go-rest-project/db"
	"github.com/hiteshnayak305/go-rest-project/internal/util"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) CreateUser() error {

	hashedpass, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)`

	stmt, err := db.SqlConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, hashedpass)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetUserByID(id int64) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE id = ?"
	row := db.SqlConnection.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE email = ?"
	row := db.SqlConnection.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
