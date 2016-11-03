package handlers

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-auth/db"
	"github.com/rafaeljesus/kyp-auth/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UsersCreate(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	password := u.Password
	u.Password = ""
	u.EncryptedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err := db.Repo.Create(u).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
