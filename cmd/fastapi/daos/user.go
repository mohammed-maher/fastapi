package daos

import "github.com/mohammed-maher/fastapi/models"

type UserDAO struct{}

func NewUserDao() *UserDAO {
	return &UserDAO{}
}

func (u *UserDAO) Find(identifier string) (*models.User, error) {
	var user models.User
	err := models.DB.Where("mobile=? OR email=?", identifier, identifier).
		Find(&user).
		Error
	return &user, err
}

func (u *UserDAO) UserExists(mobile,email string) bool{
	return models.DB.Where("mobile=? OR email=?", mobile, email).
		Find(&models.User{}).
		Error==nil
}

func (u *UserDAO) Create(user *models.User) error {
	return models.DB.Create(&user).Error
}

func (u *UserDAO) Update(user *models.User) error{
	return models.DB.Save(&user).Error
}