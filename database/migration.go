package database

import (
	"fmt"
	"landtick/models"
	"landtick/pkg/postgres"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Tiket{},
		&models.Kereta{},
		&models.Transaksi{},
	)

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("SETAN")
}
