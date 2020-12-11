package services

import (
	"errors"
	"github.com/mohammed-maher/fastapi/daos"
	"github.com/mohammed-maher/fastapi/models"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"net/http"
)

type tripDao interface {
	Find(uint64) (*models.Trip, error)
	Create(*models.Trip) error
	Delete(*models.Trip) error
	Get(pageNum int,srcCityId,dstCityId uint) (*[]models.Trip, error)
	Count(srcCityId,dstCityId uint) (int,error)
}

type TripService struct {
	dao tripDao
}

func NewTripService(dao tripDao) *TripService {
	return &TripService{dao: dao}
}

func (s *TripService) Create(request *requests.TripRequest) *response.Base {
	if err := request.Validate(false); err != nil {
		return response.ERROR(http.StatusBadRequest, err.Error())
	}
	if err := s.dao.Create(&models.Trip{
		CarID:         request.CarID,
		FromCityID:    request.FromCityID,
		ToCityID:      request.ToCityID,
		Passengers:    request.Passengers,
		TwoWay:        request.IsTwoWay,
		DepartureDate: request.DepartureDate,
		ReturnDate:    request.ReturnDate,
		StatusID:      1,
	}); err != nil {
		return response.ERROR(http.StatusInternalServerError, "unknown_error")
	}

	return response.OK("trip_created_successfully")
}

func (s *TripService) Delete(tripId, userId uint) *response.Base {
	trip, err := s.dao.Find(uint64(tripId))
		if err != nil{
			return response.ERROR(http.StatusNotFound, "trip_not_found")
		}
	carService:=NewCarService(daos.NewCarDao())
	car,err:=carService.dao.Find(trip.CarID)
	if err!=nil || car.UserID!=userId{
		return response.ERROR(http.StatusBadRequest,"unknown_error")
	}

	if err := s.dao.Delete(trip); err != nil {
		return response.ERROR(http.StatusUnprocessableEntity, "trip_not_deleted")
	}
	return response.OK("trip_deleted_successfully")
}

func (s *TripService) Get(pageNumber int,srcCityId,dstCityId uint) *response.PaginatedData{
	res:=response.PaginatedData{}
	data,err:=s.dao.Get(pageNumber,srcCityId,dstCityId)
	res.Error=err
	resultsCount,err:=s.dao.Count(srcCityId,dstCityId)
	if err!=nil{
		res.Error=errors.New("unknown error")
		return &res
	}
	resultsPerPage:=10
	totalPages := int(resultsCount / resultsPerPage)
	res.ResultsPerPage=uint(resultsPerPage)
	res.PagesCount=uint(totalPages)
	res.Items=*data

}


