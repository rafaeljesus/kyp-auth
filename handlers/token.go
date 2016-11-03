package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-auth/db"
	"github.com/rafaeljesus/kyp-auth/models"
	"golang.org/x/crypto/bcrypt"
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

	user := models.User{}
	if err := db.Repo.Where("email = ?", u.Email).Find(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if invalid, _ := verifyPassword(user, u.Password); !invalid {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	claims := &jwtCustomClaims{
		u.Email,
		0,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

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

type tokenResponse struct {
	Token string `json:"token"`
}

func verifyPassword(u models.User, password string) (bool, error) {
	return bcrypt.CompareHashAndPassword(u.EncryptedPassword, []byte(password)) == nil, nil
}
