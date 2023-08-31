package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawpassword string) ([]byte, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(rawpassword), 8)

	if err != nil {
		return nil, err
	}

	return hashPassword, nil

}

func VerifyPassword(rawpassword string, hashPassword string) bool {

	var err error
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawpassword))

	if err != nil {
		return false
	}
	return true

}
