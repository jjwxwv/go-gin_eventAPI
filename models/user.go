package models

import (
	"errors"

	"example.com/project/db"
	"example.com/project/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

//we should pass receiver as a pointer to get/manage the data from the original
func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt,err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//hash password before save
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email,hashPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

//check login credential 
//we should pass receiver as a pointer to get/manage the data from the original
func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	//get one row of user
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	//extract the data that we get back using Scan()
	err := row.Scan(&u.ID,&retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}
 
	//check password
	passwordValid := utils.CheckPassword(u.Password, retrievedPassword)
	if !passwordValid {
		return errors.New("credentials invalid")
	}

	return nil
}