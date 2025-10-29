package nodes

import (
	"eercase/models/eercase/enum"
	nodes2 "eercase/models/eercase/nodes"
	"eercase/pattern"
	"strconv"
	"strings"
)

// Inheritance: node com label e disjointness
type InheritanceDTO struct {
	NodeDTO
	Label        string                `gorm:"type:varchar(255)" json:"label"`
	Disjointness enum.DisjointnessType `gorm:"not null" json:"disjointness" validate:"required"`
}

func (l *InheritanceDTO) ToEntity() nodes2.Inheritance {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.InheritancePrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.InheritancePrefix)
	} else {
		rawID = l.ID
	}

	uid64, err := strconv.ParseUint(rawID, 10, 64)
	var id uint
	if err == nil {
		id = uint(uid64)
	} else {
		id = 0 // cria novo se a string não for numérica
	}

	return nodes2.Inheritance{
		Node: nodes2.Node{
			ID:         id,
			ProjectID:  l.ProjectID,
			PosistionX: l.PosistionX,
			PosistionY: l.PosistionY,
			Width:      l.Width,
			Height:     l.Height,
		},
		Label:        l.Label,
		Disjointness: l.Disjointness,
	}
}

func (l InheritanceDTO) GetID() string {
	return l.ID
}

func NewInheritanceDTOFromEntity(l *nodes2.Inheritance) InheritanceDTO {
	return InheritanceDTO{
		NodeDTO: NewNodeDTOFromEntity(&l.Node, pattern.InheritancePrefix),
		Label:   l.Label,
	}
}
