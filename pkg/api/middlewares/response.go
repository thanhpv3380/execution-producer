package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func WrapResponseHandler(handler func(c *fiber.Ctx) (interface{}, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := handler(c)

		if err != nil {
			return err
		}

		if data == nil {
			return c.JSON(map[string]string{"message": "success"})
		}

		return c.JSON(data)
	}
}
