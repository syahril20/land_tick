package handler

import (
	resultdto "landtick/dto/result"
	transaksidto "landtick/dto/transaksi"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerTransaksis struct {
	TransaksiRepositories repositories.TransaksiRepositories
}

func HandlerTransaksi(TransaksiRepositories repositories.TransaksiRepositories) *HandlerTransaksis {
	return &HandlerTransaksis{TransaksiRepositories}
}

func (h *HandlerTransaksis) FindTransaksi(c echo.Context) error {
	transaksi, err := h.TransaksiRepositories.FindTransaksi()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaksi})

}

func (h *HandlerTransaksis) FindTransaksiId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, _ := h.TransaksiRepositories.FindTransaksiId(id)

	if transaksi.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaksi})
}

func (h *HandlerTransaksis) DeleteTransaksi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, _ := h.TransaksiRepositories.FindTransaksiId(id)

	if transaksi.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.TransaksiRepositories.DeleteTransaksi(id, transaksi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
}

func (h *HandlerTransaksis) CreateTransaksi(c echo.Context) error {
	request := new(transaksidto.CreateTransaksi)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// transaksi, _ := h.TransaksiRepositories.FindTransaksiId(Id)

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
	Users, err := h.TransaksiRepositories.GetUserId(request.IdUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}
	Tikets, err := h.TransaksiRepositories.GetTiketId(request.IdTiket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	transaksi := models.Transaksi{
		TanggalTransaksi: request.TanggalTransaksi,
		Qty:              request.Qty,
		Total:            request.Total,
		Status:           request.Status,
		IdUser:           request.IdUser,
		User:             Users,
		IdTiket:          request.IdTiket,
		Tiket:            Tikets,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := h.TransaksiRepositories.CreateTransaksi(transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
}

func (h *HandlerTransaksis) UpdateTransaksi(c echo.Context) error {
	request := new(transaksidto.UpdateTransaksi)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, err := h.TransaksiRepositories.FindTransaksiId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.TanggalTransaksi != "" {
		transaksi.TanggalTransaksi = request.TanggalTransaksi
	}
	if request.Qty != 0 {
		transaksi.Qty = request.Qty
	}
	if request.Total != 0 {
		transaksi.Total = request.Total
	}
	if request.Status != "" {
		transaksi.Status = request.Status
	}
	if request.IdUser != 0 {
		transaksi.IdUser = request.IdUser
	}
	if request.IdTiket != 0 {
		transaksi.IdTiket = request.IdTiket
	}

	transaksi.UpdatedAt = time.Now()

	data, err := h.TransaksiRepositories.UpdateTransaksi(id, transaksi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
}

func convertResponseTransaksi(Transaksi models.Transaksi) transaksidto.TransaksiResponse {
	return transaksidto.TransaksiResponse{
		Id:               Transaksi.Id,
		TanggalTransaksi: Transaksi.TanggalTransaksi,
		Qty:              Transaksi.Qty,
		Total:            Transaksi.Total,
		Status:           Transaksi.Status,
		IdUser:           Transaksi.IdUser,
		User:             Transaksi.User,
		IdTiket:          Transaksi.IdTiket,
		Tiket:            Transaksi.Tiket,
	}
}
