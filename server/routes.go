package server

import (
	"github.com/ArifulProtik/arifulprotik-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func publicroutes(e fiber.Router) {
	e.Post("/user/signup", controllers.Signup)
	e.Post("/user/signin", controllers.Signin)

}

func router(e *fiber.App) {
	main := e.Group("/api/v1")
	publicroutes(main)
}
