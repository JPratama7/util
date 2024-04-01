package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/JPratama7/util/convert"
	"net/mail"
	"reflect"
)

type emailAlias = mail.Address

type Address struct {
	emailAlias
}

func (a Address) MarshalJSON() ([]byte, error) {
	if a.Address == "" {
		return nil, &json.MarshalerError{
			Type: reflect.TypeOf(a),
			Err:  errors.New("email: address is empty"),
		}
	}

	byteBuilder := new(bytes.Buffer)
	defer func() {
		byteBuilder.Reset()
	}()

	byteBuilder.Grow(len(a.Name) + len(a.Address) + 3)

	if a.Name != "" {
		byteBuilder.WriteString(a.Name + " ")
	}

	byteBuilder.WriteString("<" + a.Address + ">")

	return byteBuilder.Bytes(), nil
}

func (a *Address) UnmarshalJSON(b []byte) error {
	if len(b) < 1 {
		return &json.MarshalerError{
			Type: reflect.TypeOf(a),
			Err:  errors.New("email: address is empty"),
		}
	}

	e, err := mail.ParseAddress(convert.UnsafeString(b))
	if err != nil {
		return err
	}

	a.Name = e.Name
	a.Address = e.Address

	return nil
}
