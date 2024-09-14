package token

import (
	"aidanwoods.dev/go-paseto"
	"time"
)

func WithExpiration(d time.Duration) TokenArgs {
	return func(token *paseto.Token) error {
		token.SetExpiration(time.Now().Add(d))
		return nil
	}
}

func WithNotBefore(d time.Duration) TokenArgs {
	return func(token *paseto.Token) error {
		token.SetNotBefore(time.Now().Add(d))
		return nil
	}
}

func WithIssuer(s string) TokenArgs {
	return func(token *paseto.Token) error {
		token.SetIssuer(s)
		return nil
	}
}

func WithSubject(s string) TokenArgs {
	return func(token *paseto.Token) error {
		token.SetSubject(s)
		return nil
	}
}

func WithAudience(s string) TokenArgs {
	return func(token *paseto.Token) error {
		token.SetAudience(s)
		return nil
	}
}

func WithClaims[T any](key string, value T) TokenArgs {
	return func(token *paseto.Token) error {
		return token.Set(key, value)
	}
}
