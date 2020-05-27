package simplejwt

import (
	"encoding/json"
	"strings"

	b64 "encoding/base64"
)

//Header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func (h Header) toBase64() string {

	marshal, _ := json.Marshal((h))
	str := b64.StdEncoding.EncodeToString(marshal)
	return strings.ReplaceAll(str, "=", "")
}

func (h *Header) setHeader(alg string, typ string) {

	h.Alg = alg
	h.Typ = typ
}
