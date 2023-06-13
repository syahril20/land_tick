package handler

import (
	resultdto "landtick/dto/result"
	userdto "landtick/dto/user"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerUsers struct {
	UserRepositories repositories.UserRepositories
}

func HandlerUser(UserRepositories repositories.UserRepositories) *HandlerUsers {
	return &HandlerUsers{UserRepositories}
}

func (h *HandlerUsers) FindUser(c echo.Context) error {
	user, err := h.UserRepositories.FindUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: user})

}

func (h *HandlerUsers) FindUserId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.UserRepositories.FindUserId(id)

	if user.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: user})
}

func (h *HandlerUsers) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.UserRepositories.FindUserId(id)

	if user.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.UserRepositories.DeleteUser(id, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseUser(data)})
}

func (h *HandlerUsers) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUser)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// user, _ := h.UserRepositories.FindUserId(Id)

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
	Roles, err := h.UserRepositories.GetRoleId(request.IdRole)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	user := models.User{
		NamaLengkap:  request.NamaLengkap,
		Username:     request.Username,
		Email:        request.Email,
		Password:     request.Password,
		JenisKelamin: request.JenisKelamin,
		Telp:         request.Telp,
		Alamat:       request.Alamat,
		IdRole:       request.IdRole,
		RoleName:     Roles,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	data, err := h.UserRepositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseUser(data)})
}

func (h *HandlerUsers) UpdateUser(c echo.Context) error {
	request := new(userdto.UpdateUser)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepositories.FindUserId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.NamaLengkap != "" {
		user.NamaLengkap = request.NamaLengkap
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if request.JenisKelamin != "" {
		user.JenisKelamin = request.JenisKelamin
	}
	if request.Telp != "" {
		user.Telp = request.Telp
	}
	if request.Alamat != "" {
		user.Alamat = request.Alamat
	}
	if request.IdRole != 0 {
		user.IdRole = request.IdRole
	}

	user.UpdatedAt = time.Now()

	data, err := h.UserRepositories.UpdateUser(id, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseUser(data)})
}

func convertResponseUser(User models.User) userdto.UserResponse {
	return userdto.UserResponse{
		Id:           User.Id,
		NamaLengkap:  User.NamaLengkap,
		Username:     User.Username,
		Email:        User.Email,
		Password:     User.Password,
		JenisKelamin: User.JenisKelamin,
		Telp:         User.Telp,
		Alamat:       User.Alamat,
		IdRole:       User.IdRole,
		RoleName:     User.RoleName,
	}
}
