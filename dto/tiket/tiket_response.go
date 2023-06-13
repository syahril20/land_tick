package tiketdto

import "landtick/models"

type TiketResponse struct {
	Id               int                   `json:"id_tiket"`
	NamaKereta       string                `json:"nama_kereta" `
	IdKereta         int                   `json:"id_kereta" `
	Kereta           models.KeretaResponse `json:"kereta"`
	TanggalBerangkat string                `json:"tanggal_berangkat" `
	StasiunBerangkat string                `json:"stasiun_berangkat" `
	JamBerangkat     string                `json:"jam_berangkat" `
	StasiunTujuan    string                `json:"stasiun_tujuan" `
	JamTiba          string                `json:"jam_tiba" `
	Harga            int                   `json:"harga" `
	Qty              int                   `json:"qty" `
}
