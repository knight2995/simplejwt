package simplejwt

import (
	"testing"
)

func Test(t *testing.T) {

	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoi642V67O07KexIn0.K_RT2Q74UkU9PcyO0xDtOnqIQyWZvhuotjSnYJKuGeM"

	jwt := New()

	jwt.SetHeader("HS256", "JWT")

	jwt.Set("name", "덕보짱")

	jwt.SetSecretKey("duckbo")

	jwt.Sign()

	if want != jwt.Token() {
		t.Errorf("\n expected:%s \nactual:%s", want, jwt.Token())
	}
}

func RegistedClaimsTest(t *testing.T) {

}
