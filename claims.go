package simplejwt

import (
	b64 "encoding/base64"
	"encoding/json"
	"strings"
)

//Claim 실제 데이터들이 저장되는 공간
//simple 하게 구현을 위해 클레임을 나누지 않으며, 공개클레임 제외
type Claims map[string]interface{}

//그 어떠한 데이터라도 받아들일 수 있어야 한다...
//구현방법은 고민해보도록 하자

func (c Claims) toBase64() string {

	marshal, _ := json.Marshal((c))
	str := b64.StdEncoding.EncodeToString(marshal)
	return strings.ReplaceAll(str, "=", "")
}
