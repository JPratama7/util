package paseto

import "time"

type Payload[T any] struct {
	Id   string    `json:"id"`
	Exp  time.Time `json:"exp"`
	Iat  time.Time `json:"iat"`
	Nbf  time.Time `json:"nbf"`
	Data T         `json:"data"`
}

type PASETO struct {
	Public   string
	Private  string
	Duration time.Duration
}
