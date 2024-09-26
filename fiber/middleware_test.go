package fiber

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Mock utility function for testing purposes
func mockInjector(ctx *fiber.Ctx, key string, value any) {
	ctx.Locals(key, value)
}

func TestLocalToHeader(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	localName := "localKey"
	headerName := "X-Custom-Header"

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals(localName, "headerValue")
		return ctx.Next()
	})

	app.Use(LocalToHeader(localName, headerName))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, "headerValue", resp.Header.Get(headerName))
}

func TestLocalToHeader_WrongType(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	localName := "localKey"
	headerName := "X-Custom-Header"

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals(localName, 123) // Incorrect type
		return ctx.Next()
	})

	app.Use(LocalToHeader(localName, headerName))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	assert.Empty(t, resp.Header.Get(headerName))
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestInjector(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	key := "injectKey"
	value := "injectedValue"

	app.Use(Injector(key, value))

	app.Get("/", func(ctx *fiber.Ctx) error {
		val := ctx.Locals(key).(string)
		assert.Equal(t, value, val)
		return ctx.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
