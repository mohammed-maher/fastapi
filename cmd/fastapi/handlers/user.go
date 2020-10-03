package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammed-maher/fastapi/daos"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"github.com/mohammed-maher/fastapi/services"
)

func Login(ctx *fiber.Ctx) error {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.LoginUser
	ctx.BodyParser(&req)
	res := userService.Login(&req)
	return response.Send(ctx, res)
}

func Register(ctx *fiber.Ctx) error {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.RegisterUser
	ctx.BodyParser(&req)
	return response.Send(ctx, userService.Register(&req))
}

func ResetPasswordInit(ctx *fiber.Ctx) error {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.ResetPasswordInit
	ctx.BodyParser(&req)
	return response.Send(ctx, userService.ResetPasswordInit(&req))
}

func ResetPasswordVerify(ctx *fiber.Ctx) error {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.ResetPasswordVerify
	ctx.BodyParser(&req)
	return response.Send(ctx, userService.PasswordResetVerify(&req))
}

func ResetPasswordConform(ctx *fiber.Ctx) error {
	userService := services.NewUserService(daos.NewUserDao())
	var req requests.ResetPasswordConform
	ctx.BodyParser(&req)
	return response.Send(ctx, userService.PasswordResetConform(&req))
}

func Logout(ctx *fiber.Ctx) error {
	return response.Send(ctx, services.Logout(ctx.Get("Authorization")))
}

func RefreshToken(ctx *fiber.Ctx) error {
	var req requests.RefreshRequest
	ctx.BodyParser(&req)
	return response.Send(ctx, services.RefreshToken(&req))
}

func ActivateUser(ctx *fiber.Ctx) error{
	userService:=services.NewUserService(daos.NewUserDao())
	var req requests.ActivateUser
	ctx.BodyParser(&req)
	return response.Send(ctx,userService.Activate(&req))
}