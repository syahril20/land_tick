package routes

import (
	"landtick/handler"
	"landtick/pkg/middleware"
	"landtick/pkg/postgres"
	"landtick/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	UserRepository := repositories.RepositoryUser(postgres.DB)
	h := handler.HandlerUser(UserRepository)
	e.GET("/user", middleware.Auth(h.FindUser))
	e.GET("/user/:id", middleware.Auth(h.FindUserId))
	e.POST("/user", middleware.Auth(h.CreateUser))
	e.PATCH("/user/:id", middleware.Auth(h.UpdateUser))
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}
