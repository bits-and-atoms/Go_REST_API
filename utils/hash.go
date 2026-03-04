package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error){
	rel,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(rel),err
}

func CheckPassHash(password, hashpass string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashpass),[]byte(password))
	return err == nil
}