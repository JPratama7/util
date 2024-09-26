package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestInjector(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	key := "testKey"
	value := "testValue"

	app.Use(func(ctx *fiber.Ctx) error {
		Injector(ctx, key, value)
		return ctx.Next()
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		val, ok := ctx.Locals(key).(string)
		assert.True(t, ok)
		assert.Equal(t, value, val)
		return ctx.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestExtractor(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	key := "testKey"
	value := "testValue"

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals(key, value)
		return ctx.Next()
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		extractedValue, ok := Extractor[string](ctx, key)
		assert.True(t, ok)
		assert.Equal(t, value, extractedValue)
		return ctx.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestExtractor_WrongType(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	key := "testKey"
	value := 123 // Integer value

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals(key, value)
		return ctx.Next()
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		extractedValue, ok := Extractor[string](ctx, key) // Attempt to extract as string
		assert.False(t, ok)
		assert.Zero(t, extractedValue) // Should be zero value for string
		return ctx.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
