package Test

import "gorm.io/gorm"

type Repository interface {
	Get(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryTest(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get(user User) (User, error) {
	err := r.db.Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
