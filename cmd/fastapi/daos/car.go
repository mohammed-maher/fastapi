package daos

import "github.com/mohammed-maher/fastapi/models"

type carDao struct{}

func NewCarDao() *carDao {
	return &carDao{}
}

func (d *carDao) Create(car *models.Car) error {
	return models.DB.Create(car).Error
}

func (d *carDao) Find(id uint) (*models.Car, error) {
	var car models.Car
	return &car, models.DB.Find(&car, id).Error
}

func (d *carDao) Delete(id uint) error {
	return models.DB.Delete(models.Car{}, id).Error
}
