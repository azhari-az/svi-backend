package database

import (
	"fmt"
	"svi-backend/repository"

	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(&repository.Post{})
	fmt.Println("Table Posts Created.")
}
