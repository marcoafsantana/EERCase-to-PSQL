package nodes

import (
	nodes2 "eercase/models/eercase/nodes"
	"strconv"

	"github.com/google/uuid"
)

// Node: Todo node deve estar atrelado a um projeto
type NodeDTO struct {
	ID         string    `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	ProjectID  uuid.UUID `gorm:"primaryKey;type:char(36);not null" json:"project_id" validate:"required,uuid4"`
	PosistionX float32   `json:"position_x"`
	PosistionY float32   `json:"position_y"`
	Width      int       `json:"width"`
	Height     int       `json:"height"`
}

func NewNodeDTOFromEntity(l *nodes2.Node, prefix string) NodeDTO {
	return NodeDTO{
		ID:         prefix + strconv.FormatUint(uint64(l.ID), 10),
		ProjectID:  l.ProjectID,
		PosistionX: l.PosistionX,
		PosistionY: l.PosistionY,
		Width:      l.Width,
		Height:     l.Height,
	}
}
