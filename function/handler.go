package function

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func verifyHashFromBytes(input []byte) error {
	bss := bytes.SplitN(input, []byte(" "), 2)

	hash := bss[0]
	pass := bss[1]

	return bcrypt.CompareHashAndPassword(hash, pass)
}

// Handle a serverless request
func Handle(input []byte) string {
	if err := verifyHashFromBytes(input); err != nil {
		return fmt.Sprintf("false %s", err.Error())
	}

	return fmt.Sprint("true")
}
