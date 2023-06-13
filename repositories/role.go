package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type RoleRepositories interface {
	FindRole() ([]models.Role, error)
	FindRoleId(Id int) (models.Role, error)
	DeleteRole(Id int, Role models.Role) (models.Role, error)
	CreateRole(Role models.Role) (models.Role, error)
	UpdateRole(Id int, Role models.Role) (models.Role, error)
}

func RepositoryRole(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindRole() ([]models.Role, error) {
	var Role []models.Role

	err := r.db.Find(&Role).Error
	return Role, err
}

func (r *repositories) FindRoleId(Id int) (models.Role, error) {
	var Roles models.Role
	err := r.db.First(&Roles, Id).Error

	return Roles, err
}

func (r *repositories) DeleteRole(Id int, Role models.Role) (models.Role, error) {
	err := r.db.Delete(&Role).Error

	return Role, err
}

func (r *repositories) CreateRole(Role models.Role) (models.Role, error) {
	err := r.db.Create(&Role).Error

	return Role, err
}

func (r *repositories) UpdateRole(Id int, Role models.Role) (models.Role, error) {
	err := r.db.Save(&Role).Error

	return Role, err
}
