package wsc_models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id       string `gorm:"size:50;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
