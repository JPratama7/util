package fiber

import "github.com/gofiber/fiber/v2"

type ReturnData[T any] struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
}

func (rcv ReturnData[T]) WriteResponseBody(ctx *fiber.Ctx) error {
	return ctx.Status(rcv.Code).JSON(rcv)
}

func NewReturnData[T any](code int, success bool, status string, data T) ReturnData[T] {
	return ReturnData[T]{
		Code:    code,
		Success: success,
		Status:  status,
		Data:    data,
	}
}
