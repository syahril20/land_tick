package models

import "time"

type Transaksi struct {
	Id               int           `json:"id_transaksi" form:"id_transaksi" gorm:"primary_key:auto_increment"`
	TanggalTransaksi string        `json:"tanggal_transaksi" form:"tanggal_transaksi" gorm:"type:varchar(20)"`
	Qty              int           `json:"qty" form:"qty"`
	Total            int           `json:"total" form:"total"`
	Status           string        `json:"status" form:"status" gorm:"type: varchar(10)"`
	IdUser           int           `json:"id_user" form:"id_user"`
	User             UserResponse  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdUser"`
	IdTiket          int           `json:"id_tiket" form:"id_tiket"`
	Tiket            TiketResponse `json:"tiket" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdTiket"`
	CreatedAt        time.Time     `json:"-"`
	UpdatedAt        time.Time     `json:"-"`
}

type TransaksiResponse struct {
	Id               int    `json:"id_transaksi"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
}

func (TransaksiResponse) TableName() string {
	return "transaksi"
}
