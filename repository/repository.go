package repository

import (
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db}
}
