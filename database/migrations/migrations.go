package migrations

import (
	"api-go.com/mod/models"
	"gorm.io/gorm"
)

// Cria tabela no banco de acordo com o Model
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Tarefa{})
}