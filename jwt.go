package simplejwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type JWT struct {
	header Header
	claims Claims
	sign   Signature
	secret string // secret Key
	token  string
}

func New() *JWT {
	jwt := JWT{}
	jwt.claims = make(Claims)
	return &jwt
}

//interface
func (jwt *JWT) SetHeader(alg string, typ string) {
	jwt.header.setHeader(alg, typ)

}
func (jwt *JWT) Set(key string, val interface{}) {

	fmt.Println(key, val)
	jwt.claims[key] = val

}

//make toekn
func (jwt *JWT) Sign() {

	jwt.secret = "duckbo"

	data := jwt.header.toBase64() + "." + jwt.claims.toBase64()

	fmt.Println(data)

	h := hmac.New(sha256.New, []byte(jwt.secret))

	h.Write([]byte(data))

	sha := hex.EncodeToString(h.Sum(nil))

	jwt.token = data + "." + sha

}

func (jwt JWT) GetToken() string {

	return jwt.token

}

//hs256

func Hello() string {
	return "Hello, world."
}
