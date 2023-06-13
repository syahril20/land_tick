package tiketdto

import "time"

type CreateTiket struct {
	NamaKereta       string    `json:"nama_kereta" form:"nama_kereta" validate:"required"`
	IdKereta         int       `json:"id_kereta" form:"id_kereta" validate:"required"`
	TanggalBerangkat string    `json:"tanggal_berangkat" form:"tanggal_berangkat" validate:"required"`
	StasiunBerangkat string    `json:"stasiun_berangkat" validate:"required"`
	JamBerangkat     string    `json:"jam_berangkat" form:"jam_berangkat" validate:"required"`
	StasiunTujuan    string    `json:"stasiun_tujuan" validate:"required"`
	JamTiba          string    `json:"jam_tiba" form:"jam_tiba" validate:"required"`
	Harga            int       `json:"harga" form:"harga" validate:"required"`
	Qty              int       `json:"qty" form:"qty" validate:"required"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

type UpdateTiket struct {
	NamaKereta       string    `json:"nama_kereta" form:"nama_kereta" `
	IdKereta         int       `json:"id_kereta" form:"id_kereta" `
	TanggalBerangkat string    `json:"tanggal_berangkat" form:"tanggal_berangkat" `
	StasiunBerangkat string    `json:"stasiun_berangkat" `
	JamBerangkat     string    `json:"jam_berangkat" form:"jam_berangkat" `
	StasiunTujuan    string    `json:"stasiun_tujuan" `
	JamTiba          string    `json:"jam_tiba" form:"jam_tiba" `
	Harga            int       `json:"harga" form:"harga" `
	Qty              int       `json:"qty" form:"qty" `
	UpdatedAt        time.Time `json:"-"`
}
