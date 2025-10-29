package nodes

import (
	nodes2 "eercase/models/eercase/nodes"
	"eercase/pattern"
	"strconv"
	"strings"
)

// AssociativeEntity: extende Entity
type AssociativeEntityDTO struct {
	EntityDTO
	RelationshipID string `gorm:"not null" json:"relationship_id"`
}

func (l *AssociativeEntityDTO) ToEntity(nodeMap map[string]string) nodes2.AssociativeEntity {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.AssociativeEntityPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.AssociativeEntityPrefix)
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

	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var relationshipRawID string
	if strings.HasPrefix(l.RelationshipID, pattern.RelationshipPrefix) {
		relationshipRawID = strings.TrimPrefix(l.RelationshipID, pattern.RelationshipPrefix)
	} else {
		relationshipRawID = l.RelationshipID
	}

	relationshipID64, err := strconv.ParseUint(relationshipRawID, 10, 64)
	var relationshipID *uint
	if err == nil {
		relationshipID = new(uint)
		*relationshipID = uint(relationshipID64)
	} else {
		if sourceIDFromList, ok := nodeMap[l.RelationshipID]; ok {
			relationshipRawID := strings.TrimPrefix(sourceIDFromList, pattern.RelationshipPrefix)
			relationshipID64, _ := strconv.ParseUint(relationshipRawID, 10, 64)
			relationshipID = new(uint)
			*relationshipID = uint(relationshipID64)
		} else {
			relationshipID = nil
		}
	}

	return nodes2.AssociativeEntity{
		Entity: nodes2.Entity{
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
		},
		RelationshipID: relationshipID,
	}
}

func NewAssociativeEntityDTOFromEntity(l *nodes2.AssociativeEntity) AssociativeEntityDTO {
	ae := AssociativeEntityDTO{
		EntityDTO: EntityDTO{
			ElementDTO: NewElementDTOFromEntity(&l.Element, pattern.AssociativeEntityPrefix),
			IsWeak:     l.IsWeak,
		},
	}

	if l.RelationshipID == nil {
		ae.RelationshipID = ""
	} else {
		ae.RelationshipID = pattern.RelationshipPrefix + strconv.FormatUint(uint64(*l.RelationshipID), 10)
	}

	return ae
}

func (l AssociativeEntityDTO) GetID() string {
	return l.ID
}
