package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (a *AuthMiddleware[ID, T]) Decode(ctx *fiber.Ctx) (err error) {
	token := ctx.Get(a.TokenHeader)
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("Missing %s header", a.TokenHeader))
	}

	decoded, err := a.token.Decode(ID(token))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Error When Decoding token")
	}

	ctx.Locals(a.TokenHeader, decoded)
	err = ctx.Next()
	return
}

func (a *AuthMiddleware[ID, T]) DataLocals(ctx *fiber.Ctx) (data T, ok bool) {
	dataI := ctx.Locals(a.TokenHeader)
	if dataI == nil {
		return
	}

	data, ok = dataI.(T)
	return
}

func (a *AuthMiddleware[ID, T]) GetID(ctx *fiber.Ctx) (data ID, ok bool) {
	token := ctx.Get(a.TokenHeader)
	if token == "" {
		return
	}

	data, err := a.token.GetId(ID(token))
	if err != nil {
		return
	}

	ok = true
	return
}
