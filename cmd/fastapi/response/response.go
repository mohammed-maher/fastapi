package response

import "github.com/gofiber/fiber"

type Response interface {
	data() map[string]interface{}
	code() int
	error() error
	message() string
}

func Send(ctx *fiber.Ctx, r Response) {
	response := fiber.Map{}
	response["success"] = true
	data:=r.data()
	if data != nil {
		response["data"]=data
	}
	message:=r.message()
	if message!=""{
		response["message"]=message
	}
	if r.error() != nil {
		response["success"] = false
		response["error"] = r.error().Error()
	}
	ctx.Status(r.code()).JSON(response)
}


