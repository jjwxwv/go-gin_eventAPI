package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//hash password by using GenerateFromPassword from bcrypt
	//need to pass a password in byte type and cost number(how complex the hashing will be)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}