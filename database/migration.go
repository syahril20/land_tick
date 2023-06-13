package database

import (
	"fmt"
	"landtick/models"
	"landtick/pkg/postgres"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Failed")
}
