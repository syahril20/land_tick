package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TransaksiRepositories interface {
	FindTransaksi() ([]models.Transaksi, error)
	FindTransaksiId(Id int) (models.Transaksi, error)
	GetTransByUser(Id int) ([]models.Transaksi, error)
	GetUserId(Id int) (models.UserResponse, error)
	GetTiketId(Id int) (models.TiketResponse, error)
	DeleteTransaksi(Id int, Transaksi models.Transaksi) (models.Transaksi, error)
	CreateTransaksi(Transaksi models.Transaksi) (models.Transaksi, error)
	UpdateTransaksi(Status string, Id int) (models.Transaksi, error)
}

func RepositoryTransaksi(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) GetTransByUser(Id int) ([]models.Transaksi, error) {

	var Transaksi []models.Transaksi
	err := r.db.Where("id_user = ?", Id).Preload("User").Find(&Transaksi).Error

	return Transaksi, err
}

func (r *repositories) FindTransaksi() ([]models.Transaksi, error) {
	var Transaksi []models.Transaksi

	err := r.db.Preload("Tiket.Kereta").Preload("User.RoleName").Find(&Transaksi).Error
	return Transaksi, err
}

func (r *repositories) FindTransaksiId(Id int) (models.Transaksi, error) {
	var Transaksis models.Transaksi
	err := r.db.Preload("Tiket.Kereta").Preload("User.RoleName").First(&Transaksis, Id).Error

	return Transaksis, err
}

func (r *repositories) GetUserId(Id int) (models.UserResponse, error) {
	var Users models.UserResponse
	err := r.db.Preload("RoleName").Preload("User").First(&Users, Id).Error

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

// func (r *repositories) CreateTransaksi(Transaksi models.Transaksi) (models.Transaksi, error) {
// 	err := r.db.Create(&Transaksi).Error

// 	return Transaksi, err
// }

// func (r *repositories) UpdateTransaksi(Id int, Transaksi models.Transaksi) (models.Transaksi, error) {
// 	err := r.db.Save(&Transaksi).Error

// 	return Transaksi, err
// }

func (r *repositories) CreateTransaksi(Transaksi models.Transaksi) (models.Transaksi, error) {
	err := r.db.Create(&Transaksi).Error

	return Transaksi, err
}

// func (r *repositories) UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
// 	err := r.db.Save(&Transaction).Error

// 	return Transaction, err
// }

func (r *repositories) UpdateTransaksi(status string, Id int) (models.Transaksi, error) {
	var transaction models.Transaksi
	r.db.Preload("Tiket").Preload("User").First(&transaction, Id)

	// if status != transaction.Status && status == "success" {
	// 	var Trip models.Trip
	// 	r.db.First(&Trip, transaction.Trip.Id)
	// 	Trip.Current_Quota = Trip.Current_Quota + transaction.CounterQty
	// 	r.db.Save(&Trip)
	// }

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
