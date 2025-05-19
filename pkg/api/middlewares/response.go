package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thanhpv3380/api/errors"
	"github.com/thanhpv3380/api/logger"
	"github.com/thanhpv3380/api/types"
)

func WrapResponseHandler(handler func(c *fiber.Ctx) (interface{}, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := handler(c)

		if err != nil {
			if customErr, ok := err.(*types.ErrorResponse); ok {
				if customErr.Message != "" {
					customErr.Message = customErr.GetMessage()
				}

				return c.Status(customErr.GetStatusCode()).JSON(customErr)
			}

			txId := c.Locals("txId").(string)
			logger.Error(fmt.Sprintf("[%s] An error occurred", txId), err)

			return c.Status(fiber.StatusInternalServerError).JSON(errors.NewError("", ""))
		}

		if data == nil {
			return c.JSON(map[string]string{"message": "success"})
		}

		return c.JSON(data)
	}
}
