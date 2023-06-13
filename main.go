package main

import (
	"fmt"
	"net/http"

	"landtick/database"
	"landtick/pkg/postgres"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	e := echo.New()

	postgres.DatabaseConnection()
	database.RunMigration()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ANJAYYani")
	})
	// PORT := os.Getenv("DB_PORT")

	fmt.Println("Server Telah Berjalan di 😘: 5000")
	e.Logger.Fatal(e.Start(":5000"))

}
