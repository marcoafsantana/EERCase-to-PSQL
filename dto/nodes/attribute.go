package nodes

import (
	enum2 "eercase/models/eercase/enum"
	nodes2 "eercase/models/eercase/nodes"
	"eercase/pattern"
	"strconv"
	"strings"
)

// Attribute: extende Element e utiliza enums do pacote enum
type AttributeDTO struct {
	ElementDTO
	Type         enum2.AttributeType   `gorm:"not null" json:"type" validate:"required"`
	DataType     enum2.DataType        `gorm:"not null" json:"data_type" validate:"required"`
	Size         float64               `gorm:"not null" json:"size" validate:"required"`
	IsNull       bool                  `gorm:"not null" json:"is_null" validate:"required"`
	IsUnique     bool                  `gorm:"not null" json:"is_unique" validate:"required"`
	Check        string                `gorm:"type:text" json:"check"`
	DefaultValue string                `gorm:"type:text" json:"default_value"`
	Comment      string                `gorm:"type:text" json:"comment"`
	Cardinality  enum2.CardinalityType `gorm:"not null" json:"cardinality" validate:"required"`
}

func (l AttributeDTO) GetID() string {
	return l.ID
}

func (l *AttributeDTO) ToEntity() nodes2.Attribute {

	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.AttributePrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.AttributePrefix)
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

	return nodes2.Attribute{
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
		Type:         l.Type,
		DataType:     l.DataType,
		Size:         l.Size,
		IsNull:       l.IsNull,
		IsUnique:     l.IsUnique,
		Check:        l.Check,
		DefaultValue: l.DefaultValue,
		Comment:      l.Comment,
		Cardinality:  l.Cardinality,
	}
}

func NewAttributeDTOFromEntity(l *nodes2.Attribute) AttributeDTO {
	return AttributeDTO{
		ElementDTO:   NewElementDTOFromEntity(&l.Element, pattern.AttributePrefix),
		Type:         l.Type,
		DataType:     l.DataType,
		Size:         l.Size,
		IsNull:       l.IsNull,
		IsUnique:     l.IsUnique,
		Check:        l.Check,
		DefaultValue: l.DefaultValue,
		Comment:      l.Comment,
		Cardinality:  l.Cardinality,
	}
}
