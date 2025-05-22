package server

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thanhpv3380/api/errors"
	"github.com/thanhpv3380/api/logger"
	"github.com/thanhpv3380/api/types"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewServer(port int) *fiber.App {
	app := fiber.New()

	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")

		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		txId := uuid.New().String()
		c.Locals("txId", txId)

		startTime := time.Now()
		c.Locals("startTime", startTime)

		logger.Infof("[%s] [REQ] %s %s", txId, c.Method(), c.OriginalURL())

		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Next()

		txId := c.Locals("txId").(string)

		startTime := c.Locals("startTime").(time.Time)
		duration := time.Since(startTime).Milliseconds()

		if c.Response().StatusCode() == fiber.StatusOK {
			logger.Infof("[%s] [RES] %d %s | duration=%dms", txId, c.Response().StatusCode(), c.OriginalURL(), duration)
		} else {
			logger.Infof("[%s] [RES] %d %s | duration=%dms | error=%s", txId, c.Response().StatusCode(), c.OriginalURL(), duration, c.Response().Body())
		}

		return nil
	})

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()

		txId := c.Locals("txId").(string)

		if err != nil {
			if fiberErr, ok := err.(*fiber.Error); ok {
				if fiberErr.Code == fiber.StatusNotFound || fiberErr.Code == fiber.StatusMethodNotAllowed {
					jsonData, _ := json.Marshal(errors.NewUriNotFound())
					return c.Status(fiber.StatusNotFound).Send(jsonData)
				}
			}

			if customErr, ok := err.(*types.ErrorResponse); ok {
				if customErr.Message == "" {
					customErr.Message = customErr.GetMessage()
				}

				return c.Status(customErr.GetStatusCode()).JSON(customErr)
			}

			logger.Error(fmt.Sprintf("[%s] An error occurred", txId), err)
			return c.Status(fiber.StatusInternalServerError).JSON(errors.NewError("", ""))
		}

		return nil
	})

	return app
}
