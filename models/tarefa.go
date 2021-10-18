package models

import (
	"time"

	// "gorm.io/gorm"
)

// Struct do tipo Tarefa
type Tarefa struct {
	ID        int64     `json:"id" gorm:"primaryKey";autoIncrement:true`
	Descricao string    `json:"descricao"`
	Prazo     time.Time `json:"prazo"`
	Completa  bool      `json:"completa"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
