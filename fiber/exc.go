package fiber

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	gjson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"log"
	"strings"
)

func CustomErrorHandler(defErrHan func(err error) error) func(*fiber.Ctx, error) error {
	if defErrHan == nil {
		return func(ctx *fiber.Ctx, err error) error {
			response := ReturnData[*string]{
				Code:    fiber.StatusInternalServerError,
				Success: false,
				Data:    nil,
				Status:  "Internal Server Error",
			}

			switch e := err.(type) {
			case *fiber.Error:
				response.Code = e.Code
				response.Status = e.Message
			case validator.ValidationErrors:
				strFinal := "Error terjadi pada field :"
				for _, v := range e {
					strFinal += fmt.Sprintf(" %s", v.Field())
				}
				response.Code = fiber.StatusBadRequest
				response.Status = strFinal
			case *gjson.InvalidUnmarshalError, *gjson.UnmarshalTypeError, *gjson.MarshalerError, *gjson.UnsupportedTypeError, *gjson.UnsupportedValueError, *gjson.SyntaxError, *gjson.PathError:
				response.Code = fiber.StatusBadRequest
				response.Status = "Invalid JSON"
			default:
				response.Code = fiber.StatusInternalServerError
				response.Status = "Internal Server Error"
			}
			return response.WriteResponseBody(ctx)
		}
	}
	return func(ctx *fiber.Ctx, err error) error {
		response := ReturnData[*string]{
			Code:    fiber.StatusInternalServerError,
			Success: false,
			Data:    nil,
			Status:  "Internal Server Error",
		}

		switch e := err.(type) {
		case *fiber.Error:
			response.Code = e.Code
			response.Status = e.Message
		case validator.ValidationErrors:
			strFinal := "Error terjadi pada field :"
			for _, v := range e {
				strFinal += fmt.Sprintf(" %s", v.Field())
			}
			response.Code = fiber.StatusBadRequest
			response.Status = strFinal
		case *gjson.InvalidUnmarshalError, *gjson.UnmarshalTypeError, *gjson.MarshalerError, *gjson.UnsupportedTypeError, *gjson.UnsupportedValueError, *gjson.SyntaxError, *gjson.PathError:
			response.Code = fiber.StatusBadRequest
			response.Status = "Invalid JSON"
		default:
			return defErrHan(err)
		}
		return response.WriteResponseBody(ctx)
	}
}

func ErrHandler(ctx *fiber.Ctx, err error) error {
	response := ReturnData[*string]{
		Code:    fiber.StatusInternalServerError,
		Success: false,
		Data:    nil,
		Status:  "Internal Server Error",
	}
	vb := strings.Builder{}
	vb.WriteString("Error when Validating")

	switch e := err.(type) {
	case *fiber.Error:
		response.Code = e.Code
		response.Status = e.Message
	case validate.Errors:
		for k, _ := range e {
			vb.WriteString(fmt.Sprintf(" %s", k))
		}
		response.Code = fiber.StatusBadRequest
		response.Status = vb.String()
	case validator.ValidationErrors:
		for _, v := range e {
			vb.WriteString(fmt.Sprintf(" %s", v.Field()))
		}
		response.Code = fiber.StatusBadRequest
		response.Status = vb.String()
	case *validator.InvalidValidationError:
		response.Code = fiber.StatusBadRequest
		response.Status = e.Error()
	case *gjson.InvalidUnmarshalError, *gjson.UnmarshalTypeError, *gjson.MarshalerError, *gjson.UnsupportedTypeError, *gjson.UnsupportedValueError, *gjson.SyntaxError, *gjson.PathError:
		response.Code = fiber.StatusBadRequest
		response.Status = "Invalid JSON"
	default:
		response.Code = fiber.StatusInternalServerError
		response.Status = "Internal Server Error"
		log.Printf("\nInternal Error : %+v\n", err)
	}
	return response.WriteResponseBody(ctx)
}
