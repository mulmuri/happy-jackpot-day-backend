package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)



func CheckPassword(existingHash, incomingPass string) error {
	if existingHash == incomingPass {
		return nil
	}
	return errors.New("pass not match")
	//return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass))
}

func HashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	sBytes := []byte(*s)
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	*s = string(hashedBytes[:])
	return nil
}