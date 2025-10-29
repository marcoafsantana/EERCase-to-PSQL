package nodes

import (
	nodes2 "eercase/models/eercase/nodes"
	"eercase/pattern"
	"strconv"
	"strings"
)

// Entity: extende Element
type EntityDTO struct {
	ElementDTO
	IsWeak bool `gorm:"not null" json:"is_weak" validate:"required"`
}

func (l *EntityDTO) ToEntity() nodes2.Entity {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.EntityPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.EntityPrefix)
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

	return nodes2.Entity{
		Element: nodes2.Element{
			Node: nodes2.Node{
				ID:         id,
				ProjectID:  l.ProjectID,
				PosistionX: l.PosistionX,
				PosistionY: l.PosistionY,
				Width:      l.Width,
				Height:     l.Height,
			},
			Name: l.Name,
		},
		IsWeak: l.IsWeak,
	}
}

func (l EntityDTO) GetID() string {
	return l.ID
}

func NewEntityDTOFromEntity(l *nodes2.Entity) EntityDTO {
	return EntityDTO{
		ElementDTO: NewElementDTOFromEntity(&l.Element, pattern.EntityPrefix),
		IsWeak:     l.IsWeak,
	}
}
