package router

import (
	"github.com/gofiber/fiber/v2"
	"go-chat/server/internal/user"
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

	app.Post("/signup", userHandler.CreateUser)
}

func Start(addr string) error {
	return app.Listen(addr)
}
