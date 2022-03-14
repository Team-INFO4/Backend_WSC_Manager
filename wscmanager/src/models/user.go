package wsc_models

import (
	"html"

	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"wscmanager.com/utils/token"
)

type User struct {
	//gorm.Model
	Id       string `gorm:"size:50;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func (u *User) SaveUser() (*User, error) {
	var err error
	DB.AutoMigrate(&User{})
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Id = html.EscapeString(u.Id)
	u.Password = string(hashedPassword)
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(id string, password string) (string, error) {
	var err error

	u := User{}
	err = DB.Model(User{}).Where("Id = ?", id).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(u.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.CreateJWT(u.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
