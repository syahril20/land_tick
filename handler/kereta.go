package handler

import (
	keretadto "landtick/dto/kereta"
	resultdto "landtick/dto/result"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerKeretas struct {
	KeretaRepositories repositories.KeretaRepositories
}

func HandlerKereta(KeretaRepositories repositories.KeretaRepositories) *HandlerKeretas {
	return &HandlerKeretas{KeretaRepositories}
}

func (h *HandlerKeretas) FindKereta(c echo.Context) error {
	kereta, err := h.KeretaRepositories.FindKereta()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: kereta})

}

func (h *HandlerKeretas) FindKeretaId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	kereta, _ := h.KeretaRepositories.FindKeretaId(id)

	if kereta.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: kereta})
}

func (h *HandlerKeretas) DeleteKereta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	kereta, _ := h.KeretaRepositories.FindKeretaId(id)

	if kereta.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.KeretaRepositories.DeleteKereta(id, kereta)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseKereta(data)})
}

func (h *HandlerKeretas) CreateKereta(c echo.Context) error {
	request := new(keretadto.CreateKereta)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// kereta, _ := h.KeretaRepositories.FindKeretaId(Id)

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

	kereta := models.Kereta{
		JenisKereta: request.JenisKereta,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	data, err := h.KeretaRepositories.CreateKereta(kereta)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseKereta(data)})
}

func (h *HandlerKeretas) UpdateKereta(c echo.Context) error {
	request := new(keretadto.UpdateKereta)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	kereta, err := h.KeretaRepositories.FindKeretaId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.JenisKereta != "" {
		kereta.JenisKereta = request.JenisKereta
	}

	kereta.UpdatedAt = time.Now()

	data, err := h.KeretaRepositories.UpdateKereta(id, kereta)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseKereta(data)})
}

func convertResponseKereta(Kereta models.Kereta) keretadto.KeretaResponse {
	return keretadto.KeretaResponse{
		Id:          Kereta.Id,
		JenisKereta: Kereta.JenisKereta,
	}
}
