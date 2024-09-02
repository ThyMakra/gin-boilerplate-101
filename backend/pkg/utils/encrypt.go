package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pwByte := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pwByte, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(result)
}

func ComparePassword(hashedPassword string, password string) error {
	pw := []byte(password)
	hashedPw := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPw, pw)
	return err
}
