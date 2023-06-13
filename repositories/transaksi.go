package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TransaksiRepositories interface {
	FindTransaksi() ([]models.Transaksi, error)
	FindTransaksiId(Id int) (models.Transaksi, error)
	GetUserId(Id int) (models.UserResponse, error)
	GetTiketId(Id int) (models.TiketResponse, error)
	DeleteTransaksi(Id int, Transaksi models.Transaksi) (models.Transaksi, error)
	CreateTransaksi(Transaksi models.Transaksi) (models.Transaksi, error)
	UpdateTransaksi(Id int, Transaksi models.Transaksi) (models.Transaksi, error)
}

func RepositoryTransaksi(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindTransaksi() ([]models.Transaksi, error) {
	var Transaksi []models.Transaksi

	err := r.db.Preload("User.RoleName").Preload("Tiket.Kereta").Find(&Transaksi).Error
	return Transaksi, err
}

func (r *repositories) FindTransaksiId(Id int) (models.Transaksi, error) {
	var Transaksis models.Transaksi
	err := r.db.Preload("User.RoleName").Preload("Tiket.Kereta").First(&Transaksis, Id).Error

	return Transaksis, err
}

func (r *repositories) GetUserId(Id int) (models.UserResponse, error) {
	var Users models.UserResponse
	err := r.db.Preload("RoleName").First(&Users, Id).Error

	return Users, err
}

func (r *repositories) GetTiketId(Id int) (models.TiketResponse, error) {
	var Tikets models.TiketResponse
	err := r.db.Preload("Kereta").First(&Tikets, Id).Error

	return Tikets, err
}

func (r *repositories) DeleteTransaksi(Id int, Transaksi models.Transaksi) (models.Transaksi, error) {
	err := r.db.Preload("User.RoleName").Preload("Tiket.Kereta").Delete(&Transaksi).Error

	return Transaksi, err
}

func (r *repositories) CreateTransaksi(Transaksi models.Transaksi) (models.Transaksi, error) {
	err := r.db.Create(&Transaksi).Error

	return Transaksi, err
}

func (r *repositories) UpdateTransaksi(Id int, Transaksi models.Transaksi) (models.Transaksi, error) {
	err := r.db.Save(&Transaksi).Error

	return Transaksi, err
}
