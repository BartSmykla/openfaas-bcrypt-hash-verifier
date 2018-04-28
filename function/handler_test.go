package function

import (
	"bytes"
	"testing"
)

func Test_verifyHashFromBytes(t *testing.T) {
	table := []struct {
		match bool
		desc  string
		input []byte
	}{
		// match
		{
			match: true,
			desc:  "simple password without numbers special characters",
			input: []byte("$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K foo"),
		},
		{
			match: true,
			desc:  "simple password with numbers and without special characters",
			input: []byte("$2a$12$/iu.578C1vnNdWZM7l2T6uTT.xG6drlHes/B2TeN/sqNdhA.dEUGe foo123"),
		},
		{
			match: true,
			desc:  "password with space inside",
			input: []byte("$2a$12$Sw30vJt8yeBU9rwtFq54oucGDQ4m02G9FLcgTz2zX0RsTf8vCmjwK foo bar"),
		},
		{
			match: true,
			desc:  "password with space at the end",
			input: []byte("$2a$12$s6Q3W3JGyysZFJaMMC62M.HZnqLuuomflNTBrWZXLGJW6dpO5A3o. foobar "),
		},
		{
			match: true,
			desc:  "password with numbers and space at the end",
			input: []byte("$2a$04$/SlUYRAzWmHYQKBIvdxCueqe68cOsUJ/C5.vN1.QKD0VjGz7ATCDW foobar123 "),
		},
		{
			match: true,
			desc:  "password with space character in the middle and at the end",
			input: []byte("$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i foo bar "),
		},
		{
			match: true,
			desc:  "password with special characters and numbers",
			input: []byte("$2a$12$51LuhvD/ChvwHItK32tG/OQSw6V36Vuq9Z8t21lcXRBXNkw4SveMK foobarbaz!@#$%^&*()123"),
		},
		{
			match: true,
			desc:  "password with special characters, numbers and space in the middle",
			input: []byte("$2a$12$mGNsnW1s.vwiqfzyojDRXeMQvr5XGyKNCxyv6DrtlB5Gt6fqVlpl2 foobarbaz !@#$%^&*()123"),
		},
		{
			match: true,
			desc:  "password with special characters, numbers, space in the middle and in the end",
			input: []byte("$2a$04$oZObgBGHknNxhSX8N/0dgOxvp/.HpRQuLt6YjDY/vl8NUdmfoqri2 foobarbaz !@#$%^&*()123 "),
		},

		// not match

		{
			match: false,
			desc:  "no hash and password",
			input: []byte(""),
		},
		{
			match: false,
			desc:  "hash present but no password",
			input: []byte("$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K"),
		},
		{
			match: false,
			desc:  "hash doesn't match password",
			input: []byte("$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K foo "),
		},
		{
			match: false,
			desc:  "hash doesn't match password",
			input: []byte("$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i foobar "),
		},
		{
			match: false,
			desc:  "hash doesn't match password",
			input: []byte("$2a$12$fWH.UrG.U2iX9FqSY8oEdOWHlzxwgaRRfl57X5.MgaFrOVUqgu12K foobar123!@##@%$^%$$@)(*^&%#$@"),
		},
		{
			match: false,
			desc:  "hash is wrong because of space inside",
			input: []byte("$2a$1 2$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i foo"),
		},
	}

	for _, c := range table {
		if c.match {
			t.Run(c.desc, func(t *testing.T) {
				if err := verifyHashFromBytes(c.input); err != nil {
					bss := bytes.SplitN(c.input, []byte(" "), 2)

					hash := bss[0]
					pass := bss[1]

					t.Errorf(
						"Should not return error in case of %s (Hash: %s) (Password: %s). Hash should match password",
						c.desc,
						hash,
						pass,
					)
				}
			})
		} else {
			t.Run(c.desc, func(t *testing.T) {
				if err := verifyHashFromBytes(c.input); err == nil {
					bss := bytes.SplitN(c.input, []byte(" "), 2)

					if len(bss) >= 2 {
						hash := bss[0]
						pass := bss[1]

						t.Errorf(
							"Should return error in case of %s (Hash: %s) (Password: %s). Hash should match password",
							c.desc,
							hash,
							pass,
						)
					} else {
						t.Errorf("Should return error in case of %s", c.desc)
					}
				}
			})
		}
	}
}
