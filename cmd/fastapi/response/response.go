package response

import "github.com/gofiber/fiber/v2"

type Response interface {
	data() map[string]interface{}
	code() int
	error() error
	message() string
}

func Send(ctx *fiber.Ctx, r Response) error {
	response := fiber.Map{}
	response["success"] = r.error() == nil
	data := r.data()
	if data != nil {
		response["data"] = data
	}
	message := r.message()
	if message != "" {
		response["message"] = message
	}
	if r.error() != nil {
		response["error"] = r.error().Error()
	}
	return ctx.Status(r.code()).JSON(response)
}
