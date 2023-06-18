package transaksidto

import "time"

type CreateTransaksi struct {
	TanggalTransaksi string    `json:"tanggal_transaksi" form:"tanggal_transaksi" validate:"required"`
	QtyDewasa        int       `json:"qty_dewasa" form:"qty_dewasa" validate:"required"`
	QtyAnak          int       `json:"qty_anak" form:"qty_anak"`
	PulangPergi      bool      `json:"pulang_pergi" form:"pulang_pergi"`
	Total            int       `json:"total" form:"total" validate:"required"`
	Status           string    `json:"status" form:"status" validate:"required"`
	IdUser           int       `json:"id_user" form:"id_user" validate:"required"`
	IdTiket          int       `json:"id_tiket" form:"id_tiket" validate:"required"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

type UpdateTransaksi struct {
	TanggalTransaksi string    `json:"tanggal_transaksi" form:"tanggal_transaksi"`
	Qty              int       `json:"qty" form:"qty"`
	Total            int       `json:"total" form:"total"`
	Status           string    `json:"status" form:"status"`
	IdUser           int       `json:"id_user" form:"id_user"`
	IdTiket          int       `json:"id_tiket" form:"id_tiket"`
	UpdatedAt        time.Time `json:"-"`
}
