package handler

import (
	resultdto "landtick/dto/result"
	tiketdto "landtick/dto/tiket"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerTikets struct {
	TiketRepositories repositories.TiketRepositories
}

func HandlerTiket(TiketRepositories repositories.TiketRepositories) *HandlerTikets {
	return &HandlerTikets{TiketRepositories}
}

func (h *HandlerTikets) FindTiket(c echo.Context) error {
	tiket, err := h.TiketRepositories.FindTiket()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: tiket})

}

func (h *HandlerTikets) FindTiketId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tiket, _ := h.TiketRepositories.FindTiketId(id)

	if tiket.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: tiket})
}

func (h *HandlerTikets) DeleteTiket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tiket, _ := h.TiketRepositories.FindTiketId(id)

	if tiket.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.TiketRepositories.DeleteTiket(id, tiket)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTiket(data)})
}

func (h *HandlerTikets) CreateTiket(c echo.Context) error {
	request := new(tiketdto.CreateTiket)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// tiket, _ := h.TiketRepositories.FindTiketId(Id)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}
	Roles, err := h.TiketRepositories.GetKeretaId(request.IdKereta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	tiket := models.Tiket{
		NamaKereta:       request.NamaKereta,
		IdKereta:         request.IdKereta,
		Kereta:           Roles,
		TanggalBerangkat: request.TanggalBerangkat,
		StasiunBerangkat: request.StasiunBerangkat,
		JamBerangkat:     request.JamBerangkat,
		StasiunTujuan:    request.StasiunTujuan,
		JamTiba:          request.JamTiba,
		Harga:            request.Harga,
		Qty:              request.Qty,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	data, err := h.TiketRepositories.CreateTiket(tiket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTiket(data)})
}

func (h *HandlerTikets) UpdateTiket(c echo.Context) error {
	request := new(tiketdto.UpdateTiket)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	tiket, err := h.TiketRepositories.FindTiketId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.NamaKereta != "" {
		tiket.NamaKereta = request.NamaKereta
	}
	if request.IdKereta != 0 {
		tiket.IdKereta = request.IdKereta
	}
	if request.TanggalBerangkat != "" {
		tiket.TanggalBerangkat = request.TanggalBerangkat
	}
	if request.StasiunBerangkat != "" {
		tiket.StasiunBerangkat = request.StasiunBerangkat
	}
	if request.JamBerangkat != "" {
		tiket.JamBerangkat = request.JamBerangkat
	}
	if request.StasiunTujuan != "" {
		tiket.StasiunTujuan = request.StasiunTujuan
	}
	if request.JamTiba != "" {
		tiket.JamTiba = request.JamTiba
	}
	if request.Harga != 0 {
		tiket.Harga = request.Harga
	}
	if request.Qty != 0 {
		tiket.Qty = request.Qty
	}

	tiket.UpdatedAt = time.Now()

	data, err := h.TiketRepositories.UpdateTiket(id, tiket)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTiket(data)})
}

func convertResponseTiket(Tiket models.Tiket) tiketdto.TiketResponse {
	return tiketdto.TiketResponse{
		Id:               Tiket.Id,
		NamaKereta:       Tiket.NamaKereta,
		IdKereta:         Tiket.IdKereta,
		Kereta:           Tiket.Kereta,
		TanggalBerangkat: Tiket.TanggalBerangkat,
		StasiunBerangkat: Tiket.StasiunBerangkat,
		JamBerangkat:     Tiket.JamBerangkat,
		StasiunTujuan:    Tiket.StasiunTujuan,
		JamTiba:          Tiket.JamTiba,
		Harga:            Tiket.Harga,
		Qty:              Tiket.Qty,
	}
}
