package routes

import (
	"landtick/handler"
	"landtick/pkg/middleware"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func TransaksiRoutes(e *echo.Group) {
	TransaksiRepository := repositories.RepositoryTransaksi(postgres.DB)
	h := handler.HandlerTransaksi(TransaksiRepository)
	e.GET("/transaksi", h.FindTransaksi)
	e.GET("/transaksis", middleware.Auth(h.GetTransByUser))
	e.GET("/transaksi/:id", middleware.Auth(h.FindTransaksiId))
	e.POST("/transaksi", middleware.Auth(h.CreateTransaksi))
	// e.PATCH("/transaction/:id", middleware.Auth(h.Notification))
	e.DELETE("/transaksi/:id", middleware.Auth(h.DeleteTransaksi))
	e.POST("/notification", h.Notification)
}
