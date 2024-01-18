package auth

import (
	"github.com/JPratama7/util"
)

func NewAuthMiddleware[ID ~string, T util.IdGetter[ID]](token util.Token[ID, T], tokenHeader string) *AuthMiddleware[ID, T] {
	return &AuthMiddleware[ID, T]{
		TokenHeader: tokenHeader,
		token:       token,
	}
}
