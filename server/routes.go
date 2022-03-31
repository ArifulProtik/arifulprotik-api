package server

import (
	"github.com/ArifulProtik/arifulprotik-api/auth"
	"github.com/ArifulProtik/arifulprotik-api/config"
	"github.com/ArifulProtik/arifulprotik-api/controllers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func publicroutes(e fiber.Router) {
	e.Post("/user/signup", controllers.Signup)
	e.Post("/user/signin", controllers.Signin)

}
func privateroutes(e fiber.Router) {
	e.Get("/users", controllers.GetUsers)
}

func router(e *fiber.App) {
	main := e.Group("/api/v1")
	publicroutes(main)

	main.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.C.Jwt_Key),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: auth.JwtError,
	}))
	privateroutes(main)
}
