package api

import (
	"github.com/IgorCastilhos/go-ecommerce-app/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	err := app.Listen(config.ServerPort)
	if err != nil {
		return
	}
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
