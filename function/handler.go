package function

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Hash     string `json:"hash"`
	Pass     string `json:"pass"`
	Password string `json:"password"`
}

func (i *Data) getPassword() (string, error) {
	password := i.Password

	if password == "" {
		password = i.Pass
	}

	if password == "" {
		return "", fmt.Errorf("there is no password")
	}

	return password, nil
}

func verifyHashFromBytes(input []byte) error {
	var d Data

	err := json.Unmarshal(input, &d)

	if err != nil {
		return err
	}

	if d.Hash == "" {
		return fmt.Errorf("you didn't pass hash")
	}

	pass, err := d.getPassword()

	if err != nil {
		return fmt.Errorf("you didn't pass password")
	}

	return bcrypt.CompareHashAndPassword([]byte(d.Hash), []byte(pass))
}

// Handle a serverless request
func Handle(input []byte) string {
	if err := verifyHashFromBytes(input); err != nil {
		return fmt.Sprintf(`{"match":false,"error":"%s"}`, err.Error())
	}

	return fmt.Sprint(`{"match":true}`)
}
