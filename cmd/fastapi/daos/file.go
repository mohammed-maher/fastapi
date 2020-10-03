package daos

import "github.com/mohammed-maher/fastapi/models"

type fileDao struct{}

func NewFileDao() *fileDao {
	return &fileDao{}
}

func (dao *fileDao) Create(fileData *models.File) (uint, error) {
	f := models.DB.Create(fileData)
	if err := f.Error; err != nil {
		return 0, err
	}
	return fileData.ID, nil
}

func (dao *fileDao) Find(id uint) (*models.File, error) {
	var file models.File
	return &file, models.DB.Find(&file, id).Error
}

func (dao *fileDao) Delete(id uint) error {
	return models.DB.Delete(&models.File{}, id).Error
}
