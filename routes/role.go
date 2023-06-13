package routes

import (
	"landtick/handler"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func RoleRoutes(e *echo.Group) {
	RoleRepository := repositories.RepositoryRole(postgres.DB)
	h := handler.HandlerRole(RoleRepository)
	e.GET("/role", h.FindRole)
}
