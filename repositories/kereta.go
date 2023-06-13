package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type KeretaRepositories interface {
	FindKereta() ([]models.Kereta, error)
	FindKeretaId(Id int) (models.Kereta, error)
	DeleteKereta(Id int, Kereta models.Kereta) (models.Kereta, error)
	CreateKereta(Kereta models.Kereta) (models.Kereta, error)
	UpdateKereta(Id int, Kereta models.Kereta) (models.Kereta, error)
}

func RepositoryKereta(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindKereta() ([]models.Kereta, error) {
	var Kereta []models.Kereta

	err := r.db.Find(&Kereta).Error
	return Kereta, err
}

func (r *repositories) FindKeretaId(Id int) (models.Kereta, error) {
	var Keretas models.Kereta
	err := r.db.First(&Keretas, Id).Error

	return Keretas, err
}

func (r *repositories) DeleteKereta(Id int, Kereta models.Kereta) (models.Kereta, error) {
	err := r.db.Delete(&Kereta).Error

	return Kereta, err
}

func (r *repositories) CreateKereta(Kereta models.Kereta) (models.Kereta, error) {
	err := r.db.Create(&Kereta).Error

	return Kereta, err
}

func (r *repositories) UpdateKereta(Id int, Kereta models.Kereta) (models.Kereta, error) {
	err := r.db.Save(&Kereta).Error

	return Kereta, err
}
