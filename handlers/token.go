package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-auth/resources"
	"net/http"
	"os"
	"time"
)

const UNAUTHORIZED = "Unauthorized"

var KYP_SECRET_KEY = os.Getenv("KYP_SECRET_KEY")

func TokenCreate(c echo.Context) error {
	u := resources.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	if u.Email == "" {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	user := resources.User{}
	if err := user.Authenticate(u.Email, u.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	claims := newJwtCustomClaims(u.Email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(KYP_SECRET_KEY))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
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
