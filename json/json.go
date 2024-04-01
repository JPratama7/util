package json

import (
	"github.com/goccy/go-json"
	"net/http"
)

type ReturnData[T any] struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
}

func NewReturnData[T any](code int, success bool, status string, data T) ReturnData[T] {
	return ReturnData[T]{
		Code:    code,
		Success: success,
		Status:  status,
		Data:    data,
	}
}

func (rcv ReturnData[T]) WriteResponseBody(ctx http.ResponseWriter) error {
	payload, err := json.Marshal(rcv)
	if err != nil {
		return err
	}

	ctx.Header().Set("Content-Type", "application/json")
	ctx.WriteHeader(rcv.Code)
	_, err = ctx.Write(payload)
	return err
}
