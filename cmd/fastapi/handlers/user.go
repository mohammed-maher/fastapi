package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/mohammed-maher/fastapi/daos"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"github.com/mohammed-maher/fastapi/services"
)

func Login(ctx *fiber.Ctx) {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.LoginUser
	ctx.BodyParser(&req)
	res := userService.Login(&req)
	response.Send(ctx, res)
}

func Register(ctx *fiber.Ctx) {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.RegisterUser
	ctx.BodyParser(&req)
	response.Send(ctx,userService.Register(&req))
}

func ResetPasswordInit(ctx *fiber.Ctx){
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.ResetPassword
	ctx.BodyParser(&req)
	response.Send(ctx,userService.ResetPasswordInit(&req))
}

func ResetPasswordConform(ctx *fiber.Ctx){
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.ResetPasswordConform
	ctx.BodyParser(&req)
	response.Send(ctx,userService.PasswordResetConform(&req))
}
