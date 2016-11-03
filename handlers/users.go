package handlers

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-auth/models"
	"net/http"
)

func UsersCreate(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	if err := user.Create().Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
