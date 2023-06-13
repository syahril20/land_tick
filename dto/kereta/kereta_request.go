package keretadto

type CreateKereta struct {
	JenisKereta string `json:"jenis_kereta" form:"jenis_kereta" validate:"required"`
}

type UpdateKereta struct {
	JenisKereta string `json:"jenis_kereta" form:"jenis_kereta"`
}
