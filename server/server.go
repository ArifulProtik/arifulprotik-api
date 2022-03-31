package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ArifulProtik/arifulprotik-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func startServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}
		log.Println("Server Is Shutting Down")

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(config.C.Port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func RunServer() {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       fmt.Sprintf("ArifulProtik.xyz %s", config.C.Version),
		StrictRouting: true,
		ServerHeader:  "ArifulProtik.xyz",
	})
	app.Use(logger.New())
	//app.Use(func() fiber.Handler {
	//	return func(c *fiber.Ctx) error {
	//		c.Accepts("json", "text")
	//		c.Status(500).JSON(fiber.Map{
	//			"Error": "Provide Fuck",
	//		})
	//		return nil
	//	}
	//})
	router(app)
	startServerWithGracefulShutdown(app)
	//log.Fatal(app.Listen(config.C.Port))
}
