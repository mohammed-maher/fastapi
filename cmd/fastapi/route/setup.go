package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mohammed-maher/fastapi/handlers"
	"github.com/mohammed-maher/fastapi/middleware"
)

func Register(app *fiber.App) {
	app.Use(logger.New())
	{
		api := app.Group("/api")
		//authorization endpoints
		{
			r := api.Group("/auth")
			r.Post("/login", handlers.Login)
			r.Post("/register", handlers.Register)
			r.Post("/activate",handlers.ActivateUser)
			r.Get("/logout", handlers.Logout)
			r.Post("/refresh", handlers.RefreshToken)
			{
				r = r.Group("/resetpassword")
				r.Post("/init", handlers.ResetPasswordInit)
				r.Post("/verify", handlers.ResetPasswordVerify)
				r.Post("/conform", handlers.ResetPasswordConform)
			}
		}

		//car endpoints
		{
			r := api.Group("/cars")
			r.Use(middleware.AuthorizeUser)
			r.Post("/add", handlers.AddCar)
			r.Post("/delete/:id", handlers.DeleteCar)
		}
	}
}
