package function

import (
	"fmt"
	"testing"
)

func Test_verifyHashFromBytes(t *testing.T) {
	table := []struct {
		shouldMatch bool
		desc        string
		hash        string
		pass        string
		format      string
	}{
		// match
		{
			shouldMatch: true,
			desc:        "simple password without numbers or special characters",
			hash:        "$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K",
			pass:        "foo",
		},
		{
			shouldMatch: true,
			desc:        "simple password with numbers and without special characters",
			hash:        "$2a$12$/iu.578C1vnNdWZM7l2T6uTT.xG6drlHes/B2TeN/sqNdhA.dEUGe",
			pass:        "foo123",
		},
		{
			shouldMatch: true,
			desc:        "password with space in the middle",
			hash:        "$2a$12$Sw30vJt8yeBU9rwtFq54oucGDQ4m02G9FLcgTz2zX0RsTf8vCmjwK",
			pass:        "foo bar",
		},
		{
			shouldMatch: true,
			desc:        "password with space at the end",
			hash:        "$2a$12$s6Q3W3JGyysZFJaMMC62M.HZnqLuuomflNTBrWZXLGJW6dpO5A3o.",
			pass:        "foobar ",
		},
		{
			shouldMatch: true,
			desc:        "password with numbers and space at the end",
			hash:        "$2a$04$/SlUYRAzWmHYQKBIvdxCueqe68cOsUJ/C5.vN1.QKD0VjGz7ATCDW",
			pass:        "foobar123 ",
		},
		{
			shouldMatch: true,
			desc:        "password with space in the middle and at the end",
			hash:        "$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i",
			pass:        "foo bar ",
		},
		{
			shouldMatch: true,
			desc:        "password with special characters and numbers",
			hash:        "$2a$12$51LuhvD/ChvwHItK32tG/OQSw6V36Vuq9Z8t21lcXRBXNkw4SveMK",
			pass:        "foobarbaz!@#$%^&*()123",
		},
		{
			shouldMatch: true,
			desc:        "password with special characters, numbers and space in the middle",
			hash:        "$2a$12$mGNsnW1s.vwiqfzyojDRXeMQvr5XGyKNCxyv6DrtlB5Gt6fqVlpl2",
			pass:        "foobarbaz !@#$%^&*()123",
		},
		{
			shouldMatch: true,
			desc:        "password with special characters, numbers, space in the middle and at the end",
			hash:        "$2a$04$oZObgBGHknNxhSX8N/0dgOxvp/.HpRQuLt6YjDY/vl8NUdmfoqri2",
			pass:        "foobarbaz !@#$%^&*()123 ",
		},

		// not match

		{
			shouldMatch: false,
			desc:        "malformed json",
			format:      "%s",
		},
		{
			shouldMatch: false,
			desc:        "malformed json (no comma between elements)",
			format:      `{"hash":"%s""password":"%s"}`,
		},
		{
			shouldMatch: false,
			desc:        "no hash and password",
			hash:        "",
			pass:        "",
		},
		{
			shouldMatch: false,
			desc:        "hash present but no password",
			hash:        "$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K",
			pass:        "",
		},
		{
			shouldMatch: false,
			desc:        "hash doesn't match password",
			hash:        "$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K",
			pass:        "foo ",
		},
		{
			shouldMatch: false,
			desc:        "hash doesn't match password",
			hash:        "$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i",
			pass:        "foobar ",
		},
		{
			shouldMatch: false,
			desc:        "hash doesn't match password",
			hash:        "$2a$12$fWH.UrG.U2iX9FqSY8oEdOWHlzxwgaRRfl57X5.MgaFrOVUqgu12K",
			pass:        "foobar123!@##@%$^%$$@)(*^&%#$@",
		},
		{
			shouldMatch: false,
			desc:        "hash is wrong because of space inside",
			hash:        "$2a$1 2$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i",
			pass:        "foo",
		},
	}

	errShouldMatch := "Should not return error in case of %s (Hash: %s) (Password: %s). Hash should match password"
	errShouldNotMatch := "Should return error in case of %s (Hash: %s) (Password: %s). Hash should not match password"

	for _, c := range table {
		defaultFormat := `{"hash":"%s","password":"%s"}`
		format := c.format

		if format == "" {
			format = defaultFormat
		}

		input := []byte(fmt.Sprintf(format, c.hash, c.pass))

		t.Run(c.desc, func(t *testing.T) {
			if err := verifyHashFromBytes(input); c.shouldMatch && err != nil {
				t.Errorf(errShouldMatch, c.desc, c.hash, c.pass)
			} else if !c.shouldMatch && err == nil {
				t.Errorf(errShouldNotMatch, c.desc, c.hash, c.pass)
			}
		})
	}
}
