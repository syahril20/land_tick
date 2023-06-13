package userdto

import "landtick/models"

type UserResponse struct {
	Id           int                 `json:"id_user"`
	NamaLengkap  string              `json:"nama_lengkap" form:"nama_lengkap"`
	Username     string              `json:"username" form:"username"`
	Email        string              `json:"email" form:"email"`
	Password     string              `json:"password" form:"password"`
	JenisKelamin string              `json:"jenis_kelamin" form:"jenis_kelamin"`
	Telp         string              `json:"telp" form:"telp"`
	Alamat       string              `json:"alamat" form:"alamat"`
	IdRole       int                 `json:"id_role" form:"id_role"`
	RoleName     models.RoleResponse `json:"role_name"`
}
