package middleware

import (
	"github.com/JPratama7/util"
	"gorm.io/gorm"
)

type AuthMiddleware[ID ~string, T util.IdGetter[ID]] struct {
	db          *gorm.DB
	TokenHeader string
	token       util.Token[ID, T]
}
