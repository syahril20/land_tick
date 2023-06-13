package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	RoleRoutes(e)
	KeretaRoutes(e)
	UserRoutes(e)
	TiketRoutes(e)
	TransaksiRoutes(e)

}
