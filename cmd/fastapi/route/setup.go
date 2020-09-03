package route

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/mohammed-maher/fastapi/handlers"
)

func Register(app *fiber.App){
	api:=app.Group("/api",middleware.Logger())
	{
		api=app.Group("/auth")
		api.Post("/login", handlers.Login)
		api.Post("/register",handlers.Register)
		api.Get("/logout",handlers.Logout)
		api.Post("/refresh",handlers.RefreshToken)
		{
			api=app.Group("/resetpassword")
			api.Post("/init", handlers.ResetPasswordInit)
			api.Post("/conform", handlers.ResetPasswordConform)
		}
	}
}
