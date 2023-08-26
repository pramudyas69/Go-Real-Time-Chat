package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-chat/server/internal/user"
	"time"
)

var app *fiber.App

func InitRouter(userHandler *user.Handler) {
	app = fiber.New()

	//app.Use(cors.New(cors.Config{
	//	AllowOrigins:     "http://localhost:3000",
	//	AllowMethods:     "GET,POST",
	//	AllowHeaders:     "Content-Type",
	//	ExposeHeaders:    "Content-Length",
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))
	// Custom middleware to log requests
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next() // Continue to next middleware
		elapsed := time.Since(start)

		fmt.Printf("Request: %s - %s | Status: %d | Duration: %v\n",
			c.Method(), c.Path(), c.Response().StatusCode(), elapsed)
		return err
	})

	app.Post("/signup", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)
}

func Start(addr string) error {
	return app.Listen(addr)
}
