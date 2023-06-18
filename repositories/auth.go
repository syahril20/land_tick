package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(auth models.User) (models.User, error)
	Login(username string) (models.User, error)
	Auth(id int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) Auth(Id int) (models.User, error) {
	var Users models.User
	err := r.db.First(&Users, Id).Error

	return Users, err
}

func (r *repositories) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repositories) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error

	return user, err
}
