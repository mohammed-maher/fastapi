package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammed-maher/fastapi/auth"
	_ "github.com/mohammed-maher/fastapi/config"
	"github.com/mohammed-maher/fastapi/models"
	"github.com/mohammed-maher/fastapi/route"
)

func main() {

	models.Setup()
	auth.SetupRedis()
	r := fiber.New(fiber.Config{
		BodyLimit: 10<<20,
	})
	route.Register(r)
	r.Listen(":8080")
}
