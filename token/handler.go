package token

import (
	"github.com/whatsauth/watoken"
)

func (r *Generator[ID, T]) Create(data T) (token string, err error) {
	token, err = watoken.EncodeWithStructDuration(string(data.GetId()), &data, r.private, r.duration)
	return
}

func (r *Generator[ID, T]) Decode(token string) (data T, err error) {
	payload, err := watoken.DecodeWithStruct[T](r.public, token)
	if err != nil {
		return
	}

	data = payload.Data
	return
}

func (r *Generator[ID, T]) FullDecode(token string) (data watoken.Payload[T], err error) {
	data, err = watoken.DecodeWithStruct[T](r.public, token)
	return
}

func (r *Generator[ID, T]) GetId(token string) (id ID, err error) {
	payload, err := watoken.Decode(r.public, token)
	if err != nil {
		return
	}

	id = ID(payload.Id)
	return

}
