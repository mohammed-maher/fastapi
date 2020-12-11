package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type PaginatedData struct {
	Base
	TotalResults   uint `json:"total_results"`
	CurrentPage    uint `json:"current_page"`
	PagesCount     uint `json:"pages_count"`
	NextPage       uint `json:"next_page"`
	PreviousPage   uint `json:"previous_page"`
	ResultsPerPage uint `json:"results_per_page"`
	Items          []interface{}
	Error          error
}

func (r *PaginatedData) Send(ctx *fiber.Ctx) error {
	if r.Error!=nil{
		return ctx.Status(http.StatusUnprocessableEntity).JSON(r.Error.Error())
	}
	response := make(map[string]interface{})
	response["total_results"] = r.TotalResults
	response["current_page"] = r.CurrentPage
	response["pages_count"] = r.PagesCount
	response["next_page"] = r.NextPage
	response["previous_page"] = r.PreviousPage
	response["results_per_page"] = r.ResultsPerPage
	response["data"] = r.Items
	return ctx.Status(http.StatusOK).JSON(response)
}
