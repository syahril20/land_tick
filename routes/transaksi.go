package routes

import (
	"landtick/handler"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func TransaksiRoutes(e *echo.Group) {
	TransaksiRepository := repositories.RepositoryTransaksi(postgres.DB)
	h := handler.HandlerTransaksi(TransaksiRepository)
	e.GET("/transaksi", h.FindTransaksi)
	e.GET("/transaksi/:id", h.FindTransaksiId)
	e.POST("/transaksi", h.CreateTransaksi)
	e.PATCH("/transaksi/:id", h.UpdateTransaksi)
	e.DELETE("/transaksi/:id", h.DeleteTransaksi)
}
