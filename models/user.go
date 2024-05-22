package models

import (
	"errors"

	"example.com/rest-api/database"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) Login() error {
	query := `
		SELECT id, password FROM users WHERE email = ?
	`

	row := database.DB.QueryRow(query, u.Email)

	var hashedPassword string

	err := row.Scan(&u.ID, &hashedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, hashedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
