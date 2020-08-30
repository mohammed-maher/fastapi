package route

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/mohammed-maher/fastapi/handlers"
)

func Register(app *fiber.App){
	api:=app.Group("/api",middleware.Logger())
	api.Post("/login", handlers.Login)
	api.Post("/register",handlers.Register)
	api.Post("/reset/init", handlers.ResetPasswordInit)
	api.Post("/reset/conform", handlers.ResetPasswordConform)
}
