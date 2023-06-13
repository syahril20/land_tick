package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type RoleRepositories interface {
	FindRole() ([]models.Role, error)
}

func RepositoryRole(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindRole() ([]models.Role, error) {
	var Role []models.Role

	err := r.db.Find(&Role).Error
	return Role, err
}
