package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

func (u *User) Crypt() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) RegistrationHandler() error {
	u.Crypt()

	DatabaseInitialization()

	db.Create(&u)

	return nil
}

var UsersSlice []User

func AddUser(user User) error {
	for _, u := range UsersSlice {
		if u.Username == user.Username {
			return errors.New("user with this nickname already exists")
		}
	}

	UsersSlice = append(UsersSlice, user)
	return nil
}
