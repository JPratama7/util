package middleware

import (
	"github.com/JPratama7/util/fiber"
	"gorm.io/gorm"
)

type AuthMiddleware[ID ~string, T fiber.IdGetter[ID]] struct {
	db          *gorm.DB
	TokenHeader string
	token       fiber.Token[ID, T]
}
