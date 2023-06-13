package userdto

import "time"

type CreateUser struct {
	NamaLengkap  string    `json:"nama_lengkap" form:"nama_lengkap" validate:"required"`
	Username     string    `json:"username" form:"username" validate:"required"`
	Email        string    `json:"email" form:"email" validate:"required"`
	Password     string    `json:"password" form:"password" validate:"required"`
	JenisKelamin string    `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
	Telp         string    `json:"telp" form:"telp" validate:"required"`
	Alamat       string    `json:"alamat" form:"alamat" validate:"required"`
	IdRole       int       `json:"id_role" form:"id_role" validate:"required"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type UpdateUser struct {
	NamaLengkap  string    `json:"nama_lengkap" form:"nama_lengkap"`
	Username     string    `json:"username" form:"username"`
	Email        string    `json:"email" form:"email"`
	Password     string    `json:"password" form:"password"`
	JenisKelamin string    `json:"jenis_kelamin" form:"jenis_kelamin"`
	Telp         string    `json:"telp" form:"telp"`
	Alamat       string    `json:"alamat" form:"alamat"`
	IdRole       int       `json:"id_role" form:"id_role"`
	UpdatedAt    time.Time `json:"-"`
}
