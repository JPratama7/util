package fiber

import "github.com/gofiber/fiber/v2"

type Response[T any] struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
}

func (r *Response[T]) SetCode(code int) *Response[T] {
	r.Code = code
	return r
}

func (r *Response[T]) SetSuccess(success bool) *Response[T] {
	r.Success = success
	return r
}

func (r *Response[T]) SetStatus(status string) *Response[T] {
	r.Status = status
	return r
}

func (r *Response[T]) SetData(data T) *Response[T] {
	r.Data = data
	return r
}

func (r *Response[T]) Write(ctx *fiber.Ctx) error {
	return ctx.Status(r.Code).JSON(r)
}

func NewResponse[T any](code int, success bool, status string, data T) Response[T] {
	return Response[T]{
		Code:    code,
		Success: success,
		Status:  status,
		Data:    data,
	}
}
