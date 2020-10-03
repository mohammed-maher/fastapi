package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)
import "github.com/mohammed-maher/fastapi/auth"

func AuthorizeUser(ctx *fiber.Ctx) error {
	token, err := auth.ExtractTokenMetadata(ctx.Get("Authorization"))
	if err != nil {
		println(err.Error())
		return ctx.SendStatus(http.StatusUnauthorized)
	}
	ctx.Request().Header.Add("X-USER-ID", fmt.Sprintf("%d", token.UserID))
	return ctx.Next()
}
