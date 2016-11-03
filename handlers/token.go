package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-auth/models"
	"net/http"
	"time"
)

func TokenCreate(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	if u.Email == "" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	user := &models.User{}
	if err := user.FindByEmail(u.Email).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if invalid, _ := user.VerifyPassword(u.Password); !invalid {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	claims := newJwtCustomClaims(u.Email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, &tokenResponse{Token: t})
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	Type uint   `json:"type"`
	jwt.StandardClaims
}

func newJwtCustomClaims(email string) *jwtCustomClaims {
	return &jwtCustomClaims{
		email,
		0,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
}

type tokenResponse struct {
	Token string `json:"token"`
}
