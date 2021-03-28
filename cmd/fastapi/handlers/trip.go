package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mohammed-maher/fastapi/daos"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"github.com/mohammed-maher/fastapi/services"
)

func AddTrip(ctx *fiber.Ctx) error {
	tripService:=services.NewTripService(daos.NewTripDao())
	var req requests.TripRequest
	ctx.BodyParser(&req)
	return response.Send(ctx,tripService.Create(&req))
}

//func GetTrips(ctx *fiber.Ctx) error{
//	from,_:=fmt.Printf("%d",ctx.Get("from"))
//	to,_:=fmt.Printf("%d",ctx.Get("to"))
//	tripService:=services.NewTripService(daos.NewTripDao())
//	return pag
//
//}
func DeleteTrip(ctx *fiber.Ctx) error {
	tripService:=services.NewTripService(daos.NewTripDao())
	tripId,_:=fmt.Printf("%d",ctx.Get("tripId"))
	userId,_:=fmt.Printf("%d",ctx.Get("X-USER-ID"))
	return response.Send(ctx,tripService.Delete(uint(tripId), uint(userId)))
}