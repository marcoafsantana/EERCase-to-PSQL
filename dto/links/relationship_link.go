package links

import (
	enum2 "eercase/models/eercase/enum"
	links2 "eercase/models/eercase/links"
	"eercase/pattern"
	"strconv"
	"strings"
)

// RelationshipLink: extende Link com participation, cardinality, role, isIdentifier e chosenLink
type RelationshipLinkDTO struct {
	LinkDTO
	Role          string                 `gorm:"type:varchar(255)" json:"role"`
	Participation enum2.CompletenessType `gorm:"not null" json:"participation" validate:"required"`
	Cardinality   enum2.CardinalityType  `gorm:"not null" json:"cardinality" validate:"required"`
	IsIdentifier  bool                   `gorm:"not null" json:"is_identifier" validate:"required"`
	ChosenLink    bool                   `gorm:"not null" json:"chosen_link" validate:"required"`
}

func (l *RelationshipLinkDTO) ToEntity(nodeMap map[string]string) links2.RelationshipLink {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.RelationshipLinkPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.RelationshipLinkPrefix)
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

	return links2.RelationshipLink{
		Link: links2.Link{
			ID:           id,
			ProjectID:    l.ProjectID,
			SourceID:     resolveID(l.SourceID, nodeMap),
			TargetID:     resolveID(l.TargetID, nodeMap),
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role:          l.Role,
		Participation: l.Participation,
		Cardinality:   l.Cardinality,
		IsIdentifier:  l.IsIdentifier,
	}
}

func NewRelationshipLinkDTOFromEntity(l *links2.RelationshipLink) RelationshipLinkDTO {
	return RelationshipLinkDTO{
		LinkDTO: LinkDTO{
			ID:           pattern.RelationshipLinkPrefix + strconv.FormatUint(uint64(l.ID), 10),
			ProjectID:    l.ProjectID,
			SourceID:     l.SourceID,
			TargetID:     l.TargetID,
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role:          l.Role,
		Participation: l.Participation,
		Cardinality:   l.Cardinality,
		IsIdentifier:  l.IsIdentifier,
		ChosenLink:    l.ChosenLink,
	}
}
