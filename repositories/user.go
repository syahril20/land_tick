package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type UserRepositories interface {
	FindUser() ([]models.User, error)
	FindUserId(Id int) (models.User, error)
	GetRoleId(Id int) (models.RoleResponse, error)
	DeleteUser(Id int, User models.User) (models.User, error)
	CreateUser(User models.User) (models.User, error)
	UpdateUser(Id int, User models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindUser() ([]models.User, error) {
	var User []models.User

	err := r.db.Preload("RoleName").Preload("Transaksi.Tiket").Preload("Transaksi.Tiket.Kereta").Find(&User).Error
	return User, err
}

func (r *repositories) FindUserId(Id int) (models.User, error) {
	var Users models.User
	err := r.db.Preload("RoleName").Preload("Transaksi.Tiket").Preload("Transaksi.Tiket.Kereta").First(&Users, Id).Error

	return Users, err
}

func (r *repositories) GetRoleId(Id int) (models.RoleResponse, error) {
	var Roles models.RoleResponse
	err := r.db.First(&Roles, Id).Error

	return Roles, err
}

func (r *repositories) DeleteUser(Id int, User models.User) (models.User, error) {
	err := r.db.Preload("RoleName").Delete(&User).Error

	return User, err
}

func (r *repositories) CreateUser(User models.User) (models.User, error) {
	err := r.db.Preload("RoleName").Create(&User).Error

	return User, err
}

func (r *repositories) UpdateUser(Id int, User models.User) (models.User, error) {
	err := r.db.Preload("RoleName").Save(&User).Error

	return User, err
}
