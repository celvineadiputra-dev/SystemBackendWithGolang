package Test

import "gorm.io/gorm"

type Repository interface {
	FindById() (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) FindById() (User, error)  {
	var user User
	err := r.db.Where("ID = 1").Find(&user).Error
	if err != nil{
		return user, err
	}
	return user, nil
}