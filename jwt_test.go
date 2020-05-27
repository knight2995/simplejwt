package simplejwt

import (
	"testing"
)

func TestingBASE64(t *testing.T) {
	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

	s := Header{}
	s.Alg = "HS256"
	s.Typ = "JWT"

	result := ToBase64(s)

	if result != want {
		t.Errorf("base64 = %q, want %q", result, want)
	}

}
