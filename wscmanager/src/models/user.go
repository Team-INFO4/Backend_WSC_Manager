package wsc_models

import (
	"html"

	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

func LoginCheck(id string, password string) (string, error) {

}
