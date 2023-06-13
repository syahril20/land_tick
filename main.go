package main

import (
	"fmt"

	"landtick/database"
	"landtick/pkg/postgres"
	"landtick/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	e := echo.New()

	postgres.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	// PORT := os.Getenv("DB_PORT")

	fmt.Println("Server Telah Berjalan di 😘: 5000")
	e.Logger.Fatal(e.Start(":5000"))

}
