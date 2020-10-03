package requests

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

type AddCarRequest struct {
	Mfr           string
	Model         string
	Year          uint
	LicenseNumber string
	LicenseGov    string
	UserID        uint
	CarPhoto      *multipart.FileHeader
	PlatePhoto    *multipart.FileHeader
}

func (r *AddCarRequest) AttachFiles(ctx *fiber.Ctx) error {
	carFile, err := ctx.FormFile("car_photo")
	if err != nil {
		return err
	}
	plateFile, err := ctx.FormFile("plate_photo")
	if err != nil {
		return err
	}
	r.CarPhoto = carFile
	r.PlatePhoto = plateFile
	return nil

}

func (r *AddCarRequest) Validate() error {
	fmt.Println(r)
	if len(r.Mfr) < 3 || len(r.Model) < 3 || len(r.LicenseNumber) < 3 || len(fmt.Sprintf("%d", r.Year)) < 4 || len(r.LicenseGov) < 3 {
		return errors.New("invalid_car_details")
	}
	if r.PlatePhoto == nil || r.CarPhoto == nil {
		return errors.New("invalid_car_files")
	}
	if r.PlatePhoto.Size == 0 || r.CarPhoto.Size == 0 {
		return errors.New("invalid_car_files")
	}
	allowedExt := []string{"JPG", "PNG", "HEIC"}
	if !validateFileExtension(r.PlatePhoto, allowedExt) || !validateFileExtension(r.CarPhoto, allowedExt) {
		return errors.New("file_not_supported")
	}
	return nil
}
