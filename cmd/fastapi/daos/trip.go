package daos

import "github.com/mohammed-maher/fastapi/models"

type tripDao struct {}

func Get(pageNum int) (*[]models.Trip,error){
	const RESULTS_PER_PAGE=10
	offset:=pageNum*RESULTS_PER_PAGE
	limit:=offset+RESULTS_PER_PAGE
	var results []models.Trip
	return &results,models.DB.Offset(offset).Limit(limit).Find(&results).Error
}

func Find(id uint64) (*models.Trip,error){
	var result models.Trip
	return &result,models.DB.Find(&result,id).Error
}

func Create(trip *models.Trip) error{
	return models.DB.Create(trip).Error
}

