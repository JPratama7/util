package paseto

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
	"github.com/JPratama7/util/types"
	"github.com/goccy/go-json"
	"time"
)

func (p PASETO) Encode(id string) (string, error) {
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(p.Duration))
	token.SetString("id", id)

	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(p.Private)
	return token.V4Sign(secretKey, nil), err

}

func (p PASETO) EncodeWithStruct(id string, data any, privateKey string) (string, error) {
	if !types.IsPointer(data) {
		return "", fmt.Errorf("data must be a pointer")
	}

	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(p.Duration))
	token.SetString("id", id)

	err := token.Set("data", data)
	if err != nil {
		return "", err
	}

	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err

}

func (p PASETO) EncodeWithStructDuration(id string, data any, dur ...time.Duration) (string, error) {
	duration := p.Duration
	if len(dur) > 0 {
		duration = dur[0]
	}

	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(duration))
	token.SetString("id", id)

	err := token.Set("data", data)
	if err != nil {
		return "", err
	}

	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(p.Private)
	return token.V4Sign(secretKey, nil), err

}

func (p PASETO) Decode(tokenString string, payload any) (err error) {
	if !types.IsPointer(payload) {
		return fmt.Errorf("payload must be a pointer")
	}

	pubKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(p.Public) // this wil fail if given key in an invalid format
	if err != nil {
		return
	}

	parser := paseto.NewParser()                                 // only used because this example token has expired, use NewParser() (which checks expiry by default)
	token, err := parser.ParseV4Public(pubKey, tokenString, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
	if err != nil {
		return
	}

	err = json.Unmarshal(token.ClaimsJSON(), &payload)
	return
}

func (p PASETO) RawDecode(tokenString string) (raw []byte, err error) {
	pubKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(p.Public) // this wil fail if given key in an invalid format
	if err != nil {
		return
	}

	parser := paseto.NewParser()                                 // only used because this example token has expired, use NewParser() (which checks expiry by default)
	token, err := parser.ParseV4Public(pubKey, tokenString, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
	if err != nil {
		return
	}

	raw = token.ClaimsJSON()
	return
}

func (p PASETO) DecodeGetId(tokenString string) string {
	payload := new(Payload[any])

	err := p.Decode(p.Public, tokenString)
	if err != nil {
		fmt.Println("Decode DecodeGetId : ", err)
	}
	return payload.Id
}
