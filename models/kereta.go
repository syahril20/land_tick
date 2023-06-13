package models

import "time"

type Kereta struct {
	Id          int       `json:"id_kereta" form:"id_kereta" gorm:"primary_key:auto_increment"`
	JenisKereta string    `json:"jenis_kereta" form:"jenis_kereta" gorm:"type: varchar(30)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type KeretaResponse struct {
	Id          int    `json:"id_kereta"`
	JenisKereta string `json:"jenis_kereta"`
}

func (KeretaResponse) TableName() string {
	return "kereta"
}
