package fiber

import (
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"os"
	"runtime"
	"strconv"
)

func GenerateConfig() fiber.Config {
	config := fiber.Config{
		JSONDecoder:  gojson.Unmarshal,
		JSONEncoder:  gojson.Marshal,
		ErrorHandler: ErrHandler,
	}

	if os.Getenv("DEV") == "true" {
		return config
	}

	if maxproc := os.Getenv("RUNTIME"); os.Getenv("PREFORK") == "true" {
		intmax := 2
		if maxproc != "" {
			conv, err := strconv.Atoi(maxproc)
			if err == nil {
				intmax = conv
			}
		}
		runtime.GOMAXPROCS(intmax)
		config.Prefork = true
	}
	return config
}
