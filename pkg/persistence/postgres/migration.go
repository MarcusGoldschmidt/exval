package postgres

import (
	"exval/pkg/models"
	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) error {
	return db.AutoMigrate(&models.ExpressionModel{})
}
