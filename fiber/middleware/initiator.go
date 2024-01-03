package middleware

import (
	"github.com/JPratama7/util"
	"gorm.io/gorm"
)

func NewAuthMiddleware[ID ~string, T util.IdGetter[ID]](db *gorm.DB, token util.Token[ID, T], tokenHeader string) *AuthMiddleware[ID, T] {
	return &AuthMiddleware[ID, T]{
		TokenHeader: tokenHeader,
		db:          db,
		token:       token,
	}
}
