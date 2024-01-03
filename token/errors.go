package token

import "errors"

var (
	ErrPrivateNotFound = errors.New("private key not found")
	ErrPublicNotFound  = errors.New("public key not found")
)
