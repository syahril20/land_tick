package routes

import (
	"landtick/handler"
	"landtick/pkg/middleware"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func RoleRoutes(e *echo.Group) {
	RoleRepository := repositories.RepositoryRole(postgres.DB)
	h := handler.HandlerRole(RoleRepository)
	e.GET("/role", middleware.Auth(h.FindRole))
	e.GET("/role/:id", middleware.Auth(h.FindRoleId))
	e.POST("/role", middleware.Auth(h.CreateRole))
	e.PATCH("/role/:id", middleware.Auth(h.UpdateRole))
	e.DELETE("/role/:id", middleware.Auth(h.DeleteRole))
}
