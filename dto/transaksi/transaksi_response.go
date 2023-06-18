package transaksidto

import "landtick/models"

type TransaksiResponse struct {
	Id               int                  `json:"id_transaksi" `
	TanggalTransaksi string               `json:"tanggal_transaksi"`
	QtyDewasa        int                  `json:"qty_dewasa"`
	QtyAnak          int                  `json:"qty_anak"`
	PulangPergi      bool                 `json:"pulang_pergi"`
	Total            int                  `json:"total"`
	Status           string               `json:"status"`
	IdUser           int                  `json:"id_user"`
	User             models.UserResponse  `json:"user"`
	IdTiket          int                  `json:"id_tiket"`
	Tiket            models.TiketResponse `json:"tiket"`
}
