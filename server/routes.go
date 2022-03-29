package server

import "github.com/gofiber/fiber/v2"

func publicroutes(e fiber.Router) {
	e.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Message": "Hello",
		})
	})
}

func router(e *fiber.App) {
	main := e.Group("/api/v1")
	publicroutes(main)
}
