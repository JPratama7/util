package fiber

import "github.com/gofiber/fiber/v2"

type ErrorResponse Response[any]

func (r *ErrorResponse) Write(ctx *fiber.Ctx) error {
	return ctx.Status(r.Code).JSON(r)
}

func (r *ErrorResponse) SetCode(code int) *ErrorResponse {
	r.Code = code
	return r
}

func (r *ErrorResponse) SetStatus(status string) *ErrorResponse {
	r.Status = status
	return r
}

func (r *ErrorResponse) SetData(data any) *ErrorResponse {
	r.Data = data
	return r
}
