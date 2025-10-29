package links

import (
	"eercase/models/eercase/enum"
	"eercase/pattern"
	"strconv"
)

// SpecializationLink: extende Link com role e type
type SpecializationLink struct {
	Link
	Role string            `gorm:"type:varchar(255)" json:"role"`
	Type enum.SLGLLinkType `gorm:"not null" json:"type" validate:"required"`
}

func (l *SpecializationLink) GetErrcaseID() string {
	// Retorna o ID do link como uma string numérica
	return pattern.SpecializationLinkPrefix + strconv.FormatUint(uint64(l.ID), 10)
}
