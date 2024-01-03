package token

import (
	"github.com/whatsauth/watoken"
)

func (r *Generator[ID, T]) Create(data T) (token string, err error) {
	if r.private == "" {
		err = ErrPrivateNotFound
		return
	}

	token, err = watoken.EncodeWithStructDuration(string(data.GetId()), &data, r.private, r.duration)
	return
}

func (r *Generator[ID, T]) Decode(token string) (data T, err error) {
	if r.public == "" {
		err = ErrPublicNotFound
		return
	}

	payload, err := watoken.DecodeWithStruct[T](r.public, token)
	if err != nil {
		return
	}

	data = payload.Data
	return
}

func (r *Generator[ID, T]) FullDecode(token string) (data watoken.Payload[T], err error) {
	if r.public == "" {
		err = ErrPublicNotFound
		return
	}

	data, err = watoken.DecodeWithStruct[T](r.public, token)
	return
}

func (r *Generator[ID, T]) GetId(token string) (id ID, err error) {
	if r.public == "" {
		err = ErrPublicNotFound
		return
	}

	payload, err := watoken.Decode(r.public, token)
	if err != nil {
		return
	}

	id = ID(payload.Id)
	return

}
