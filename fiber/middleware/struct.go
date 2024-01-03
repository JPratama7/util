package middleware

import (
	"github.com/JPratama7/util"
)

type AuthMiddleware[ID ~string, T util.IdGetter[ID]] struct {
	TokenHeader string
	token       util.Token[ID, T]
}
