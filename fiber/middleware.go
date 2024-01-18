package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func LocalToHeader(localName, headerName string) (h fiber.Handler) {
	return func(ctx *fiber.Ctx) (err error) {
		val, ok := ctx.Locals(localName).(string)
		if !ok {
			return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("local to header failed; wrong type? %T", val))
		}

		ctx.Set(headerName, val)
		err = ctx.Next()
		return
	}
}
