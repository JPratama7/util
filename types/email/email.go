package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/JPratama7/util/convert"
	"net/mail"
	"reflect"
)

type Address struct {
	m mail.Address
}

func FromNetMail(a mail.Address) Address {
	return Address{
		m: a,
	}
}

func ParseAddress(addr string) (*Address, error) {
	r, e := mail.ParseAddress(addr)
	if e != nil {
		return nil, e
	}

	return &Address{
		m: *r,
	}, nil
}

func (a Address) MarshalJSON() ([]byte, error) {
	if a.m.Address == "" {
		return nil, &json.MarshalerError{
			Type: reflect.TypeOf(a),
			Err:  errors.New("email: address is empty"),
		}
	}

	if _, err := mail.ParseAddress(a.m.Address); err != nil {
		return nil, err
	}

	byteBuilder := new(bytes.Buffer)
	defer func() {
		byteBuilder.Reset()
	}()

	byteBuilder.Grow(len(a.m.Name) + len(a.m.Address) + 3)

	if a.m.Name != "" {
		byteBuilder.WriteString(a.m.Name)
		byteBuilder.WriteRune(' ')
	}

	byteBuilder.WriteRune('<')
	byteBuilder.WriteString(a.m.Address)
	byteBuilder.WriteRune('>')

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

	a.m.Name = e.Name
	a.m.Address = e.Address

	return nil
}

func (a Address) String() string {
	return a.m.String()
}
