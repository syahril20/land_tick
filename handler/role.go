package handler

import (
	resultdto "landtick/dto/result"
	roledto "landtick/dto/role"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerRoles struct {
	RoleRepositories repositories.RoleRepositories
}

func HandlerRole(RoleRepositories repositories.RoleRepositories) *HandlerRoles {
	return &HandlerRoles{RoleRepositories}
}

func (h *HandlerRoles) FindRole(c echo.Context) error {
	role, err := h.RoleRepositories.FindRole()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: role})

}

func (h *HandlerRoles) FindRoleId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	role, _ := h.RoleRepositories.FindRoleId(id)

	if role.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: role})
}

func (h *HandlerRoles) DeleteRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	role, _ := h.RoleRepositories.FindRoleId(id)

	if role.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.RoleRepositories.DeleteRole(id, role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseRole(data)})
}

func (h *HandlerRoles) CreateRole(c echo.Context) error {
	request := new(roledto.CreateRole)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// role, _ := h.RoleRepositories.FindRoleId(Id)

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

	role := models.Role{
		NameRole:  request.NameRole,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := h.RoleRepositories.CreateRole(role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseRole(data)})
}

func (h *HandlerRoles) UpdateRole(c echo.Context) error {
	request := new(roledto.UpdateRole)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	role, err := h.RoleRepositories.FindRoleId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.NameRole != "" {
		role.NameRole = request.NameRole
	}

	role.UpdatedAt = time.Now()

	data, err := h.RoleRepositories.UpdateRole(id, role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseRole(data)})
}

func convertResponseRole(Role models.Role) roledto.RoleResponse {
	return roledto.RoleResponse{
		Id:       Role.Id,
		NameRole: Role.NameRole,
	}
}
