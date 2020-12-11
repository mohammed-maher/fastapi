package daos

import "github.com/mohammed-maher/fastapi/models"

type tripDao struct{}

func NewTripDao() *tripDao {
	return &tripDao{}
}

func (dao *tripDao) Get(pageNum int,srcCityId,dstCityId uint) (*[]models.Trip, error) {
	const RESULTS_PER_PAGE = 10
	offset := pageNum * RESULTS_PER_PAGE
	limit := offset + RESULTS_PER_PAGE
	var results []models.Trip
	return &results, models.DB.Offset(offset).Limit(limit).Where("from_city_id=? AND to_city_id=?",srcCityId,dstCityId).Find(&results).Error
}

func (dao *tripDao) Count(srcCityId,dstCityId uint) (int,error){
	var resultsCount int
	return resultsCount,models.DB.Where("from_city_id=? AND to_city_id=?",srcCityId,dstCityId).Count(&resultsCount).Error
}

func (dao *tripDao) Find(id uint64) (*models.Trip, error) {
	var result models.Trip
	return &result, models.DB.Find(&result, id).Error
}

func (dao *tripDao) Create(trip *models.Trip) error {
	return models.DB.Create(trip).Error
}

func (dao *tripDao) Delete(trip *models.Trip) error {
	return models.DB.Delete(trip).Error
}
