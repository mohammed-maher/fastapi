package services

import (
	"github.com/mohammed-maher/fastapi/helpers"
	"github.com/mohammed-maher/fastapi/models"
	"mime/multipart"
	"path/filepath"
)

type fileDao interface {
	Create(*models.File) (uint, error)
	Find(uint) (*models.File, error)
	Delete(uint) error
}
type FileService struct {
	dao fileDao
}

func NewFileService(dao fileDao) *FileService {
	return &FileService{dao: dao}
}

func (s *FileService) Upload(f *multipart.FileHeader) (uint, error) {
	obj := helpers.NewStorageObject(helpers.UPLOADS_BUCKET, f)
	if err := obj.Upload(); err != nil {
		return 0, err
	}
	return s.dao.Create(&models.File{
		Bucket: obj.Bucket,
		Object: obj.Key,
		Mime:   f.Header.Get("Content-Type"),
		Ext:    filepath.Ext(f.Filename),
	})
}

func (s *FileService) Delete(id uint) error {
	f, err := s.dao.Find(id)
	if err != nil {
		return err
	}
	obj := helpers.StorageObject{
		Key:    f.Object,
		Bucket: f.Bucket,
	}
	if err := obj.Delete(); err != nil {
		return err
	}
	return s.dao.Delete(id)
}
