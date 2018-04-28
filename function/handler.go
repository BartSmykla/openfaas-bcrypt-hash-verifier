package function

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func verifyHashFromBytes(input []byte) error {
	bss := bytes.SplitN(input, []byte(" "), 2)

	if len(bss) < 2 {
		return fmt.Errorf("you need to pass BCrypt hash and password separated by space"+
			"(for example: $2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K foo) as a parameter."+
			"You've passed: %s", string(input))
	}

	hash := bss[0]
	pass := bss[1]

	return bcrypt.CompareHashAndPassword(hash, pass)
}

// Handle a serverless request
func Handle(input []byte) string {
	if err := verifyHashFromBytes(input); err != nil {
		return fmt.Sprintf(`{"match":false,"error":"%s","passed":"%s"}`, err.Error(), string(input))
	}

	return fmt.Sprint(`{"match":true}`)
}
