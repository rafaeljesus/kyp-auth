package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/kyp-auth/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email             string `json:"email" valid:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword []byte `json:"-" sql:"encrypted_password;not null"`
}

func (u *User) Create() *gorm.DB {
	password := u.Password
	u.Password = ""
	u.EncryptedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return db.Repo.Create(u)
}

func (u *User) FindByEmail(email string) *gorm.DB {
	return db.Repo.Where("email = ?", email).Find(u)
}

func (u *User) VerifyPassword(password string) (bool, error) {
	return bcrypt.CompareHashAndPassword(u.EncryptedPassword, []byte(password)) == nil, nil
}
