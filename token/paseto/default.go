package paseto

import (
	"fmt"
	"github.com/goccy/go-json"
	"time"
)

func Encode(id string, privateKey string) (string, error) {
	return NewPASETO("", privateKey).Encode(id)
}

func EncodeWithStruct[T any](id string, data *T, privateKey string) (string, error) {
	return NewPASETO("", privateKey).EncodeWithStruct(id, data, privateKey)
}

func EncodeWithStructDuration[T any](id string, data *T, privateKey string, dur ...time.Duration) (string, error) {
	return NewPASETO("", privateKey).EncodeWithStructDuration(id, data, dur...)

}

func EncodeforHours(id string, privateKey string, hours int32) (string, error) {
	return NewPASETO("", privateKey, time.Duration(hours)*time.Hour).Encode(id)

}

func EncodeforMinutes(id string, privateKey string, minutes int32) (string, error) {
	return NewPASETO("", privateKey, time.Duration(minutes)*time.Hour).Encode(id)

}

func EncodeforSeconds(id string, privateKey string, seconds int32) (string, error) {
	return NewPASETO("", privateKey, time.Duration(seconds)*time.Hour).Encode(id)

}

func Decode[T any](publicKey string, tokenString string) (payload Payload[T], err error) {
	raw, err := NewPASETO(publicKey, "").RawDecode(tokenString)
	if err != nil {
		return
	}

	err = json.Unmarshal(raw, &payload)
	return
}
func DecodeWithStruct[T any](publicKey string, tokenstring string) (payload Payload[T], err error) {
	raw, err := NewPASETO(publicKey, "").RawDecode(tokenstring)
	if err != nil {
		return
	}

	err = json.Unmarshal(raw, &payload)
	return
}

func DecodeGetId(publicKey string, tokenstring string) string {
	payload, err := Decode[any](publicKey, tokenstring)
	if err != nil {
		fmt.Println("Decode DecodeGetId : ", err)
	}
	return payload.Id
}
