package nodes

import (
	nodes2 "eercase/models/eercase/nodes"
	"eercase/pattern"
	"strconv"
	"strings"
)

// Relationship: extende Element
type RelationshipDTO struct {
	ElementDTO
	IsIdentifier bool `gorm:"not null" json:"isIdentifier" validate:"required"`
}

func (l *RelationshipDTO) ToEntity() nodes2.Relationship {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.RelationshipPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.RelationshipPrefix)
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

	return nodes2.Relationship{
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
		IsIdentifier: l.IsIdentifier,
	}
}

func (l RelationshipDTO) GetPrefix() string {
	return pattern.RelationshipPrefix
}

func (l RelationshipDTO) GetID() string {
	return l.ID
}

func NewRelationshipDTOFromEntity(l *nodes2.Relationship) RelationshipDTO {
	return RelationshipDTO{
		ElementDTO:   NewElementDTOFromEntity(&l.Element, pattern.RelationshipPrefix),
		IsIdentifier: l.IsIdentifier,
	}
}
