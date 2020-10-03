package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammed-maher/fastapi/daos"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"github.com/mohammed-maher/fastapi/services"
	"log"
)

func AddCar(ctx *fiber.Ctx) error {
	s := services.NewCarService(daos.NewCarDao())
	var req requests.AddCarRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println(err)
	}
	if err := req.AttachFiles(ctx); err != nil {
		log.Println(err)
	}
	return response.Send(ctx, s.Add(&req))
}

func DeleteCar(ctx *fiber.Ctx) error {
	s := services.NewCarService(daos.NewCarDao())
	carId := ctx.Params("id")
	userId := ctx.Get("X-USER-ID")
	return response.Send(ctx, s.Delete(carId, userId))
}
