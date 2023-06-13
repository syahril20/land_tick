package handler

import (
	resultdto "landtick/dto/result"
	"landtick/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerRoles struct {
	RoleRepository repositories.RoleRepositories
}

func HandlerRole(RoleRepository repositories.RoleRepositories) *HandlerRoles {
	return &HandlerRoles{RoleRepository}
}

func (h HandlerRoles) FindRole(c echo.Context) error {
	role, err := h.RoleRepository.FindRole()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: role})
}
