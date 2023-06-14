package models

import "time"

type User struct {
	Id           int          `json:"id_user" form:"id_user" gorm:"primary_key:auto_increment"`
	NamaLengkap  string       `json:"nama_lengkap" form:"nama_lengkap" gorm:"type: varchar(30)"`
	Username     string       `json:"username" form:"username" gorm:"type: varchar(20)"`
	Email        string       `json:"email" form:"email" gorm:"type: varchar(50)"`
	Password     string       `json:"password" form:"password" gorm:"type: varchar(20)"`
	JenisKelamin string       `json:"jenis_kelamin" form:"jenis_kelamin" gorm:"type: varchar(10)"`
	Telp         string       `json:"telp" form:"telp" gorm:"type: varchar(15)"`
	Alamat       string       `json:"alamat" form:"alamat" gorm:"type: varchar(50)"`
	IdRole       int          `json:"id_role" form:"id_role"`
	RoleName     RoleResponse `json:"role_name" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdRole"`
	CreatedAt    time.Time    `json:"-"`
	UpdatedAt    time.Time    `json:"-"`
}

type UserResponse struct {
	Id           int          `json:"id_user"`
	NamaLengkap  string       `json:"nama_lengkap"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	JenisKelamin string       `json:"jenis_kelamin"`
	Telp         string       `json:"telp"`
	Alamat       string       `json:"alamat"`
	IdRole       int          `json:"id_role" form:"id_role"`
	RoleName     RoleResponse `json:"role_name" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdRole"`
}

func (UserResponse) TableName() string {
	return "users"
}
