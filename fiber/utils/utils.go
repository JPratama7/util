package utils

import "github.com/gofiber/fiber/v2"

func Extractor[T any](ctx *fiber.Ctx, key string) (value T, ok bool) {
	value, ok = ctx.Locals(key).(T)
	return
}

func Injector[T any](ctx *fiber.Ctx, key string, value T) {
	ctx.Locals(key, value)
}
