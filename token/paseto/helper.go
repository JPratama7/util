package paseto

import (
	"aidanwoods.dev/go-paseto"
	"time"
)

func GenerateKey() (privateKey, publicKey string) {
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey = secretKey.Public().ExportHex()     // DO share this one
	privateKey = secretKey.ExportHex()
	return privateKey, publicKey
}

func NewPASETO(publicKey, privateKey string, duration ...time.Duration) PASETO {
	dur := 2 * time.Hour
	if len(duration) > 0 {
		dur = duration[0]
	}
	return PASETO{publicKey, privateKey, dur}

}
