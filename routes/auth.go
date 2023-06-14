package routes

import (
	"landtick/handler"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(postgres.DB)
	h := handler.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}
