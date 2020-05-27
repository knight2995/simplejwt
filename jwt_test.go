package simplejwt

import (
	"testing"
)

func Test(t *testing.T) {

	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoi642V67O07KexIn0.2bf453d90ef852453d3dcc8ed310ed3a7a88432599be1ba8b634a76092ae19e3"

	jwt := New()

	jwt.SetHeader("HS256", "JWT")

	jwt.Set("name", "덕보짱")

	jwt.Sign()

	if want != jwt.GetToken() {
		t.Errorf("expected:%s actual:%s", want, jwt.GetToken())
	}
}
