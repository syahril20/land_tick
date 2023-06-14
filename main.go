package main

import (
	"fmt"
	"os"

	"landtick/database"
	"landtick/pkg/postgres"
	"landtick/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	postgres.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	PORT := os.Getenv("PORT")

	fmt.Println("Server Telah Berjalan di " + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))

}
