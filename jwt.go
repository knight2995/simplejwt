package simplejwt

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"strings"
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
func (jwt *JWT) SetSecretKey(key string) {

	jwt.secret = key
}

//make toekn
func (jwt *JWT) Sign() {

	data := jwt.header.toBase64() + "." + jwt.claims.toBase64()

	fmt.Println(data)

	h := hmac.New(sha256.New, []byte(jwt.secret))

	h.Write([]byte(data))

	//sha := hex.EncodeToString(h.Sum(nil))

	//fmt.Println((sha))
	base64Sha := strings.ReplaceAll(b64.URLEncoding.EncodeToString(h.Sum(nil)), "=", "")

	jwt.token = data + "." + base64Sha

}

/*
	jwt signed token
*/
func (jwt JWT) Token() string {

	return jwt.token

}

//hs256

func Hello() string {
	return "Hello, world."
}
