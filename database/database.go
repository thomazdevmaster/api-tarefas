package database

import (
	"log"
	"time"

	"api-go.com/mod/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// armazena a conexão
var db *gorm.DB

func StartDB() {
	// parametros para conectar no banco
	str := "host=localhost port=25432 user=admin dbname=api-tarefas password=golang"
	// str:= "DATABASE_URL"

	// Abrindo a conexão
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	// Trata possíveis erros de conexão
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	// Configurações do banco
	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	// Migrations iniciais
	migrations.RunMigrations(db)
}

// retorta a conexão
func GetDatabase() *gorm.DB {
	return db
}
