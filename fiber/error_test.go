package fiber

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestErrorResponse_Write(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	app.Get("/", func(ctx *fiber.Ctx) error {
		resp := &ErrorResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalid input",
		}
		return resp.Write(ctx)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestErrorResponse_SetCode(t *testing.T) {
	resp := &ErrorResponse{}
	resp.SetCode(404)
	assert.Equal(t, 404, resp.Code)
}

func TestErrorResponse_SetStatus(t *testing.T) {
	resp := &ErrorResponse{}
	resp.SetStatus("Not Found")
	assert.Equal(t, "Not Found", resp.Status)
}

func TestErrorResponse_SetData(t *testing.T) {
	resp := &ErrorResponse{}
	data := "Some data"
	resp.SetData(data)
	assert.Equal(t, data, resp.Data)
}
