package models

import "time"

type Tiket struct {
	Id               int            `json:"id_tiket" form:"id_tiket" gorm:"primary_key:auto_increment"`
	NamaKereta       string         `json:"nama_kereta" form:"nama_kereta" gorm:"type: varchar(30)"`
	IdKereta         int            `json:"id_kereta" form:"id_kereta"`
	Kereta           KeretaResponse `json:"jenis_kereta" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdKereta"`
	TanggalBerangkat string         `json:"tanggal_berangkat" form:"tanggal_berangkat" gorm:"type: varchar(20)"`
	StasiunBerangkat string         `json:"stasiun_berangkat" gorm:"type: varchar(20)"`
	JamBerangkat     string         `json:"jam_berangkat" form:"jam_berangkat" gorm:"type: varchar(10)"`
	StasiunTujuan    string         `json:"stasiun_tujuan" gorm:"type: varchar(20)"`
	JamTiba          string         `json:"jam_tiba" form:"jam_tiba" gorm:"type: varchar(10)"`
	Harga            int            `json:"harga" form:"harga"`
	Qty              int            `json:"qty" form:"qty"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
}

type TiketResponse struct {
	Id               int            `json:"id_tiket"`
	NamaKereta       string         `json:"nama_kereta"`
	IdKereta         int            `json:"id_kereta" form:"id_kereta"`
	Kereta           KeretaResponse `json:"jenis_kereta" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdKereta"`
	TanggalBerangkat string         `json:"tanggal_berangkat"`
	StasiunBerangkat string         `json:"stasiun_berangkat"`
	JamBerangkat     string         `json:"jam_berangkat"`
	StasiunTujuan    string         `json:"stasiun_tujuan"`
	JamTiba          string         `json:"jam_tiba"`
	Harga            int            `json:"harga" `
	Qty              int            `json:"qty" `
}

func (TiketResponse) TableName() string {
	return "tikets"
}
