package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TiketRepositories interface {
	FindTiket() ([]models.Tiket, error)
	FindTiketId(Id int) (models.Tiket, error)
	GetKeretaId(Id int) (models.KeretaResponse, error)
	DeleteTiket(Id int, Tiket models.Tiket) (models.Tiket, error)
	CreateTiket(Tiket models.Tiket) (models.Tiket, error)
	UpdateTiket(Id int, Tiket models.Tiket) (models.Tiket, error)
}

func RepositoryTiket(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindTiket() ([]models.Tiket, error) {
	var Tiket []models.Tiket

	err := r.db.Preload("Kereta").Find(&Tiket).Error
	return Tiket, err
}

func (r *repositories) FindTiketId(Id int) (models.Tiket, error) {
	var Tikets models.Tiket
	err := r.db.Preload("Kereta").First(&Tikets, Id).Error

	return Tikets, err
}

func (r *repositories) GetKeretaId(Id int) (models.KeretaResponse, error) {
	var Keretas models.KeretaResponse
	err := r.db.First(&Keretas, Id).Error

	return Keretas, err
}

func (r *repositories) DeleteTiket(Id int, Tiket models.Tiket) (models.Tiket, error) {
	err := r.db.Preload("Kereta").Delete(&Tiket).Error

	return Tiket, err
}

func (r *repositories) CreateTiket(Tiket models.Tiket) (models.Tiket, error) {
	err := r.db.Create(&Tiket).Error

	return Tiket, err
}

func (r *repositories) UpdateTiket(Id int, Tiket models.Tiket) (models.Tiket, error) {
	err := r.db.Preload("Kereta").Save(&Tiket).Error

	return Tiket, err
}
