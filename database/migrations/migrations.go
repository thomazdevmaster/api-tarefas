package migrations

import (
	"api-go.com/mod/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Tarefa{})
}