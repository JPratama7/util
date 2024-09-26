package fiber

import (
	"github.com/goccy/go-json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestResponse_SetCode(t *testing.T) {
	resp := &Response[string]{}
	resp.SetCode(200)
	assert.Equal(t, 200, resp.Code)
}

func TestResponse_SetSuccess(t *testing.T) {
	resp := &Response[string]{}
	resp.SetSuccess(true)
	assert.True(t, resp.Success)
}

func TestResponse_SetStatus(t *testing.T) {
	resp := &Response[string]{}
	resp.SetStatus("OK")
	assert.Equal(t, "OK", resp.Status)
}

func TestResponse_SetData(t *testing.T) {
	resp := &Response[string]{}
	data := "This is a test"
	resp.SetData(data)
	assert.Equal(t, data, resp.Data)
}

func TestResponse_Write(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	app.Get("/", func(ctx *fiber.Ctx) error {
		resp := &Response[string]{
			Code:    200,
			Success: true,
			Status:  "OK",
			Data:    "Hello, World!",
		}
		return resp.Write(ctx)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var responseBody Response[string]
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, 200, responseBody.Code)
	assert.True(t, responseBody.Success)
	assert.Equal(t, "OK", responseBody.Status)
	assert.Equal(t, "Hello, World!", responseBody.Data)
}

func TestNewResponse(t *testing.T) {
	code := 201
	success := true
	status := "Created"
	data := "New resource created"
	resp := NewResponse(code, success, status, data)

	assert.Equal(t, code, resp.Code)
	assert.Equal(t, success, resp.Success)
	assert.Equal(t, status, resp.Status)
	assert.Equal(t, data, resp.Data)
}
