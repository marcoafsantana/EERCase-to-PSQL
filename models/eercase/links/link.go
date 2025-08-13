package links

import (
	"github.com/google/uuid"
)

// Link: Todo link deve estar atrelado a um projeto
type Link struct {
	ID           uint      `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	FlowID       string    `gorm:"not null" json:"flow_id"`
	ProjectID    uuid.UUID `gorm:"primaryKey;type:uuid;not null" json:"project_id" validate:"required,uuid4"`
	SourceID     uint      `gorm:"not null" json:"source_id" validate:"required"`
	TargetID     uint      `gorm:"not null" json:"target_id" validate:"required"`
	SourceHandle string    `gorm:"not null" json:"source_handle" validate:"required"`
	TargetHandle string    `gorm:"not null" json:"target_handle" validate:"required"`
}
