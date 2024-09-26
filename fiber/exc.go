package fiber

import (
	"fmt"
	syncer "github.com/JPratama7/util/sync"
	"github.com/go-playground/validator/v10"
	gjson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"log"
	"strings"
	"sync"
)

var (
	oncer      sync.Once
	stringPool *syncer.Pool[*strings.Builder]
)

func initer() {
	oncer.Do(func() {
		stringPool = syncer.NewPool(func() *strings.Builder {
			return new(strings.Builder)
		})
	})
}

func NewErrorHandler(opt ...fiber.ErrorHandler) fiber.ErrorHandler {

	if len(opt) == 0 {
		return DefaultHandler
	}

	return opt[0]

}

func DefaultHandler(ctx *fiber.Ctx, err error) error {
	initer()

	response := Response[*string]{
		Code:    fiber.StatusInternalServerError,
		Success: false,
		Data:    nil,
		Status:  "Internal Server Error",
	}
	vb := stringPool.Get()
	defer func() {
		vb.Reset()
		stringPool.Put(vb)
	}()

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
	return response.Write(ctx)
}
