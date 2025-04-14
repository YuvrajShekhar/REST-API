package models

import (
	"errors"
	"example/restapi/db"
	utils "example/restapi/utlis"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

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

func (u *User) ValidateCredentails() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retreivePassword string
	err := row.Scan(&u.ID, &retreivePassword)

	if err != nil {
		return errors.New("credentials Invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retreivePassword)

	if !passwordIsValid {
		return errors.New("credentails invalid")
	}

	return nil
}
