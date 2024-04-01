package email

import (
	"github.com/stretchr/testify/assert"
	"net/mail"
	"testing"
)

func TestMarshalJSONWithValidAddress(t *testing.T) {
	addr := FromNetMail(mail.Address{
		Name:    "John Doe",
		Address: "john.doe@example.com",
	})

	data, err := addr.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "John Doe <john.doe@example.com>", string(data))
}

func TestMarshalJSONWithEmptyAddress(t *testing.T) {
	addr := FromNetMail(mail.Address{
		Name:    "John Doe",
		Address: "",
	})

	data, err := addr.MarshalJSON()
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestMarshalJSONWithEmptyName(t *testing.T) {
	addr := FromNetMail(mail.Address{
		Name:    "",
		Address: "john.doe@example.com",
	})

	data, err := addr.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "<john.doe@example.com>", string(data))
}
