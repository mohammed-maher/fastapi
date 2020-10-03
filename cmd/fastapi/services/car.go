package services

import (
	"fmt"
	"github.com/mohammed-maher/fastapi/daos"
	"github.com/mohammed-maher/fastapi/models"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"log"
	"net/http"
	"strconv"
)

type carDao interface {
	Create(*models.Car) error
	Find(uint) (*models.Car, error)
	Delete(uint) error
}
type CarService struct {
	dao carDao
}

func NewCarService(d carDao) *CarService {
	return &CarService{dao: d}
}

func (s *CarService) Add(r *requests.AddCarRequest) *response.Base {
	carInfo := models.Car{}
	if err := r.Validate(); err != nil {
		return response.ERROR(http.StatusBadRequest, err.Error())
	}
	fs := NewFileService(daos.NewFileDao())
	carPhotoId, carErr := fs.Upload(r.CarPhoto)
	platePhotoId, plateErr := fs.Upload(r.PlatePhoto)
	if carErr != nil || plateErr != nil {
		log.Println(carErr, plateErr)
		return response.ERROR(http.StatusInternalServerError, "file_upload_failed")
	}
	carInfo.PlateImageFileID = platePhotoId
	carInfo.ImageFileID = carPhotoId
	carInfo.Plate = r.LicenseNumber
	carInfo.Year = r.Year
	carInfo.UserID = r.UserID
	carInfo.Mfr = r.Mfr
	carInfo.Gov = r.LicenseGov
	carInfo.Name = r.Model
	if err := s.dao.Create(&carInfo); err != nil {
		return response.ERROR(http.StatusInternalServerError, "car_not_added")
	}
	return response.OK("car_added_successfully")
}

//Delete car
func (s *CarService) Delete(carIdParam, userIdHeader string) *response.Base {
	carId, err := strconv.ParseUint(carIdParam, 10, 64)
	if err != nil {
		return response.ERROR(http.StatusBadRequest, "incorrect_car")
	}
	userId, err := strconv.ParseUint(userIdHeader, 10, 64)
	if err != nil {
		return response.ERROR(http.StatusBadRequest, "incorrect_user")
	}
	car, err := s.dao.Find(uint(carId))
	if err != nil {
		return response.ERROR(http.StatusNotFound, "item_not_found")
	}
	if car.UserID != uint(userId) {
		fmt.Println(car.UserID,userId)
		return response.ERROR(http.StatusUnauthorized, "not_allowed")
	}
	fs := NewFileService(daos.NewFileDao())
	if err := fs.Delete(car.ImageFileID); err != nil {
		log.Println("failed to delete car image with error", err)
	}
	if err := fs.Delete(car.PlateImageFileID); err != nil {
		log.Println("failed to delete car plate with error", err)
	}
	if err := s.dao.Delete(uint(carId)); err != nil {
		return response.ERROR(http.StatusInternalServerError, "unknown_error")
	}

	return response.OK("car_deleted_successfully")
}

func (s *CarService) Update(){

}