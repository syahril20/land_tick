package routes

import (
	"landtick/handler"
	"landtick/pkg/middleware"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func TiketRoutes(e *echo.Group) {
	TiketRepository := repositories.RepositoryTiket(postgres.DB)
	h := handler.HandlerTiket(TiketRepository)
	e.GET("/tiket", h.FindTiket)
	e.GET("/tiket/:id", h.FindTiketId)
	e.POST("/tiket", middleware.Auth(h.CreateTiket))
	e.PATCH("/tiket/:id", middleware.Auth(h.UpdateTiket))
	e.DELETE("/tiket/:id", middleware.Auth(h.DeleteTiket))
}
