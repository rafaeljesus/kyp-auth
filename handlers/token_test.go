package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var KYP_USERS_ENDPOINT = os.Getenv("KYP_USERS_ENDPOINT")

func TestTokenCreate(t *testing.T) {
	defer gock.Off()
	gock.New(KYP_USERS_ENDPOINT).
		Post("/").
		Reply(http.StatusOK).
		JSON(map[string]interface{}{"id": 1, "email": "foo@mail.com"})

	body := `{"id": 1, "email": "foo@mail.com"}`
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/v1/token", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, TokenCreate(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
