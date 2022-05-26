package database

import (
	"fmt"
	"log"
	"svi-backend/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConn() *gorm.DB {
	db_host := utils.EnvVariable("DB_HOST")
	db_user := utils.EnvVariable("DB_USER")
	db_password := utils.EnvVariable("DB_PASSWORD")
	db_name := utils.EnvVariable("DB_NAME")
	db_port := utils.EnvVariable("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Koneksi Gagal")
	}
	fmt.Println("Koneksi Sukses")
	return db
}
