package routes

import (
	"landtick/handler"
	"landtick/pkg/middleware"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func KeretaRoutes(e *echo.Group) {
	KeretaRepository := repositories.RepositoryKereta(postgres.DB)
	h := handler.HandlerKereta(KeretaRepository)
	e.GET("/kereta", h.FindKereta)
	e.GET("/kereta/:id", h.FindKeretaId)
	e.POST("/kereta", middleware.Auth(h.CreateKereta))
	e.PATCH("/kereta/:id", middleware.Auth(h.UpdateKereta))
	e.DELETE("/kereta/:id", middleware.Auth(h.DeleteKereta))
}
