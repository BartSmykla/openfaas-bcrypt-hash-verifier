package function

import "testing"

func Test_verifyHashFromBytes_when_HashAndPasswordMatches(t *testing.T) {
	cases := [4][]byte{
		[]byte("$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K foo"),
		[]byte("$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i foo bar "),
		[]byte("$2a$12$fWH.UrG.U2iX9FqSY8oEdOWHlzxwgaRRfl57X5.MgaFrOVUqgu12K foo bar 123 !@##@%$^%$ $@)(*^&%#$@"),
		[]byte("$2a$12$JuDuh3z7lo89C9pCitu6jO52gIqKi3ZNCvSgWh8rwOoZIhJwF3Z2S !@#$%^&*() )(*&^%$#@!"),
	}

	for _, c := range cases {
		if err := verifyHashFromBytes(c); err != nil {
			t.Error("Should not return error. Hash match password")
		}
	}
}

func Test_verifyHashFromBytes_when_HashAndPasswordNotMatches(t *testing.T) {
	cases := [8][]byte{
		[]byte("$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K foo "),
		[]byte("$2a$12$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i foobar "),
		[]byte("$2a$12$fWH.UrG.U2iX9FqSY8oEdOWHlzxwgaRRfl57X5.MgaFrOVUqgu12K foobar123!@##@%$^%$$@)(*^&%#$@"),
		[]byte("$2a$12$JuDuh3z7lo89C9pCitu6jO52gIqKi3ZNCvSgWh8rwOoZIhJwF3Z2S !@#$%^ &*())(*& ^%$#@!"),
		[]byte("$2a$12$Y/98WmHkm3 k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K foo"),
		[]byte("$2a$1 2$EhgdZn0PRy9div6rAOBFSeNTkD9vmITuLKLD7mREEdo2tFg2Wns7i foo bar "),
		[]byte("$2a$12$fWH.UrG.U2iX9 FqSY8oEdOWHlzxwgaRRfl57X5.MgaFrOVUqgu12K foo bar 123 !@##@%$^%$ $@)(*^&%#$@"),
		[]byte("$2a$12$JuDuh3z7lo89C9pCitu6jO52gIqKi3ZNCvSgWh8r wOoZIhJwF3Z2S !@#$%^&*() )(*&^%$#@!"),
	}

	for _, c := range cases {
		if err := verifyHashFromBytes(c); err == nil {
			t.Error("Should return error. Hash not match password")
		}
	}
}
