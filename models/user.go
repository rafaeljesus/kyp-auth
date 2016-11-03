package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email             string `json:"email" valid:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword []byte `json:"-" sql:"encrypted_password;not null"`
}
