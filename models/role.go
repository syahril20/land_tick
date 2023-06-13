package models

import "time"

type Role struct {
	Id        int       `json:"id_role" form:"id_role" gorm:"primary_key:auto_increment"`
	NameRole  string    `json:"name_role" form:"name_role" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type RoleResponse struct {
	Id       int    `json:"id_role"`
	NameRole string `json:"name_role"`
}

func (RoleResponse) TableName() string {
	return "roles"
}
