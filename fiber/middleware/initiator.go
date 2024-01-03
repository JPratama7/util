package middleware

import (
	"github.com/JPratama7/util/fiber"
	"gorm.io/gorm"
)

func NewAuthMiddleware[ID ~string, T fiber.IdGetter[ID]](db *gorm.DB, token fiber.Token[ID, T], tokenHeader string) *AuthMiddleware[ID, T] {
	return &AuthMiddleware[ID, T]{
		TokenHeader: tokenHeader,
		db:          db,
		token:       token,
	}
}
