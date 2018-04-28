package function

import (
	"fmt"
	"testing"
)

func Test_verifyHashFromBytes(t *testing.T) {
	table := []struct {
		match bool
		desc  string
		hash  string
		pass  string
	}{
		// match
		{
			match: true,
			desc:  "simple password without numbers special characters",
			hash:  "$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K",
			pass:  "foo",
		},
		{
			match: true,
			desc:  "simple password with numbers and without special characters",
			hash:  "$2a$12$/iu.578C1vnNdWZM7l2T6uTT.xG6drlHes/B2TeN/sqNdhA.dEUGe",
			pass:  "foo123",
		},
		{
			match: true,
			desc:  "password with space inside",
			hash:  "$2a$12$Sw30vJt8yeBU9rwtFq54oucGDQ4m02G9FLcgTz2zX0RsTf8vCmjwK",
			pass:  "foo bar",
		},
		{
			match: true,
			desc:  "password with space at the end",
			hash:  "$2a$12$s6Q3W3JGyysZFJaMMC62M.HZnqLuuomflNTBrWZXLGJW6dpO5A3o.",
			pass:  "foobar ",
		},
		{
			match: true,
			desc:  "password with numbers and space at the end",
			hash:  "$2a$04$/SlUYRAzWmHYQKBIvdxCueqe68cOsUJ/C5.vN1.QKD0VjGz7ATCDW",
			pass:  "foobar123 ",
		},
		{
			match: true,
			desc:  "password with space character in the middle and at the end",
			hash:  "$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i",
			pass:  "foo bar ",
		},
		{
			match: true,
			desc:  "password with special characters and numbers",
			hash:  "$2a$12$51LuhvD/ChvwHItK32tG/OQSw6V36Vuq9Z8t21lcXRBXNkw4SveMK",
			pass:  "foobarbaz!@#$%^&*()123",
		},
		{
			match: true,
			desc:  "password with special characters, numbers and space in the middle",
			hash:  "$2a$12$mGNsnW1s.vwiqfzyojDRXeMQvr5XGyKNCxyv6DrtlB5Gt6fqVlpl2",
			pass:  "foobarbaz !@#$%^&*()123",
		},
		{
			match: true,
			desc:  "password with special characters, numbers, space in the middle and in the end",
			hash:  "$2a$04$oZObgBGHknNxhSX8N/0dgOxvp/.HpRQuLt6YjDY/vl8NUdmfoqri2",
			pass:  "foobarbaz !@#$%^&*()123 ",
		},

		// not match

		{
			match: false,
			desc:  "no hash and password",
			hash:  "",
			pass:  "",
		},
		{
			match: false,
			desc:  "hash present but no password",
			hash:  "$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K",
			pass:  "",
		},
		{
			match: false,
			desc:  "hash doesn't match password",
			hash:  "$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K",
			pass:  "foo ",
		},
		{
			match: false,
			desc:  "hash doesn't match password",
			hash:  "$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i",
			pass:  "foobar ",
		},
		{
			match: false,
			desc:  "hash doesn't match password",
			hash:  "$2a$12$fWH.UrG.U2iX9FqSY8oEdOWHlzxwgaRRfl57X5.MgaFrOVUqgu12K",
			pass:  "foobar123!@##@%$^%$$@)(*^&%#$@",
		},
		{
			match: false,
			desc:  "hash is wrong because of space inside",
			hash:  "$2a$1 2$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i",
			pass:  "foo",
		},
	}

	for _, c := range table {
		if c.match {
			t.Run(c.desc, func(t *testing.T) {
				input := []byte(fmt.Sprintf(`{"hash":"%s","password":"%s"}`, c.hash, c.pass))

				if err := verifyHashFromBytes(input); err != nil {
					t.Errorf(
						"Should not return error in case of %s (Hash: %s) (Password: %s). Hash should match password",
						c.desc,
						c.hash,
						c.pass,
					)
				}
			})
		} else {
			t.Run(c.desc, func(t *testing.T) {
				input := []byte(fmt.Sprintf(`{"hash":"%s","password":"%s"}`, c.hash, c.pass))

				if err := verifyHashFromBytes(input); err == nil {
					t.Errorf(
						"Should return error in case of %s (Hash: %s) (Password: %s). Hash should match password",
						c.desc,
						c.hash,
						c.pass,
					)
				}
			})
		}
	}
}
