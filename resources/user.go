package resources

import (
	"bytes"
	"encoding/json"
	"github.com/rafaeljesus/kyp-structs"
	"net/http"
	"os"
)

var KYP_USERS_ENDPOINT = os.Getenv("KYP_USERS_ENDPOINT")

type User structs.User

func (u *User) Authenticate(email string, password string) error {
	user := User{Email: email, Password: password}
	bf := new(bytes.Buffer)
	json.NewEncoder(bf).Encode(user)

	res, err := http.Post(KYP_USERS_ENDPOINT+"/authenticate", "application/json; charset=utf-8", bf)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if err != nil {
		return err
	}

	json.NewDecoder(res.Body).Decode(&u)

	return nil
}
