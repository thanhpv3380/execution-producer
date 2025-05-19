package server

import (
	"encoding/json"
	"fmt"

	"github.com/thanhpv3380/api/errors"
	"github.com/thanhpv3380/api/logger"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewServer(port int) *fiber.App {
	app := fiber.New()

	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			txId := c.Locals("txId").(string)
			logger.Error(fmt.Sprintf("[%s] An error occurred", txId), err)

			if fiberErr, ok := err.(*fiber.Error); ok {
				if fiberErr.Code == fiber.StatusNotFound || fiberErr.Code == fiber.StatusMethodNotAllowed {
					jsonData, _ := json.Marshal(errors.NewUriNotFound())
					return c.Status(fiber.StatusNotFound).Send(jsonData)
				}
			}

			return c.Status(fiber.StatusInternalServerError).JSON(errors.NewError("", ""))
		}

		return nil
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		txId := uuid.New().String()
		c.Locals("txId", txId)

		logger.Infof("[%s] [Request] %s %s", txId, c.Method(), c.OriginalURL())
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()

		txId := c.Locals("txId").(string)
		if c.Response().StatusCode() == fiber.StatusOK {
			logger.Infof("[%s] [Response] Status: %d", txId, c.Response().StatusCode())
		} else {
			logger.Infof("[%s] [Response] Status: %d %s", txId, c.Response().StatusCode(), c.Response().Body())
		}

		return err
	})

	return app
}
