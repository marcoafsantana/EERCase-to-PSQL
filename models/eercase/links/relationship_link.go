package links

import (
	enum2 "eercase/models/eercase/enum"
	"eercase/pattern"
	"strconv"
)

// RelationshipLink: extende Link com participation, cardinality, role, isIdentifier e chosenLink
type RelationshipLink struct {
	Link
	Role          string                 `gorm:"type:varchar(255)" json:"role"`
	Participation enum2.CompletenessType `gorm:"not null" json:"participation" validate:"required"`
	Cardinality   enum2.CardinalityType  `gorm:"not null" json:"cardinality" validate:"required"`
	IsIdentifier  bool                   `gorm:"not null" json:"is_identifier" validate:"required"`
	ChosenLink    bool                   `gorm:"not null" json:"chosen_link" validate:"required"`
}

func (l *RelationshipLink) GetErrcaseID() string {
	// Retorna o ID do link como uma string numérica
	return pattern.RelationshipLinkPrefix + strconv.FormatUint(uint64(l.ID), 10)
}
