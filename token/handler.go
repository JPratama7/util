package token

import (
	"github.com/JPratama7/util/token/paseto"
)

func (r *Generator[ID, T]) Create(data T) (token string, err error) {
	if r.private == "" {
		err = ErrPrivateNotFound
		return
	}

	token, err = paseto.EncodeWithStructDuration(string(data.GetId()), &data, r.private, r.duration)
	return
}

func (r *Generator[ID, T]) Decode(token string) (data T, err error) {
	if r.public == "" {
		err = ErrPublicNotFound
		return
	}

	payload, err := paseto.DecodeWithStruct[T](r.public, token)
	if err != nil {
		return
	}

	data = payload.Data
	return
}

func (r *Generator[ID, T]) FullDecode(token string) (data paseto.Payload[T], err error) {
	if r.public == "" {
		err = ErrPublicNotFound
		return
	}

	data, err = paseto.DecodeWithStruct[T](r.public, token)
	return
}

func (r *Generator[ID, T]) GetId(token string) (id ID, err error) {
	if r.public == "" {
		err = ErrPublicNotFound
		return
	}

	payload, err := paseto.Decode[T](r.public, token)
	if err != nil {
		return
	}

	id = ID(payload.Id)
	return

}
