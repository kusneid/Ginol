package user

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func (u *Credentials) Crypt() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Credentials) RegistrationHandler() bool {
	u.Crypt()
	return SendRequest(*u, os.Getenv("SERVER_REG_API_URL"))
}
