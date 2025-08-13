package nodes

import (
	"github.com/google/uuid"
)

// Node: Todo node deve estar atrelado a um projeto
type Node struct {
	ID         uint      `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	ProjectID  uuid.UUID `gorm:"primaryKey;type:uuid;not null" json:"project_id" validate:"required,uuid4"`
	PosistionX int       `json:"posistion_x"`
	PosistionY int       `json:"posistion_y"`
	Width      int       `json:"width"`
	Height     int       `json:"height"`
}
