package nodes

import (
	"github.com/google/uuid"
)

// Node: Todo node deve estar atrelado a um projeto
type Node struct {
	ID         uint      `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	ProjectID  uuid.UUID `gorm:"primaryKey;type:char(36);not null" json:"project_id" validate:"required,uuid4"`
	PosistionX float32   `json:"posistion_x"`
	PosistionY float32   `json:"posistion_y"`
	Width      int       `json:"width"`
	Height     int       `json:"height"`
}
