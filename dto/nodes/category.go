package nodes

import (
	nodes2 "eercase/models/eercase/nodes"
	"eercase/pattern"
	"strconv"
	"strings"
)

// Category: node com label
type CategoryDTO struct {
	NodeDTO
	Label string `gorm:"type:varchar(255)" json:"label"`
}

func (l *CategoryDTO) ToEntity() nodes2.Category {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.CategoryPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.CategoryPrefix)
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

	return nodes2.Category{
		Node: nodes2.Node{
			ID:         id,
			ProjectID:  l.ProjectID,
			PosistionX: l.PosistionX,
			PosistionY: l.PosistionY,
			Width:      l.Width,
			Height:     l.Height,
		},
		Label: l.Label,
	}
}

func (l CategoryDTO) GetID() string {
	return l.ID
}

func NewCategoryDTOFromEntity(l *nodes2.Category) CategoryDTO {
	return CategoryDTO{
		NodeDTO: NewNodeDTOFromEntity(&l.Node, pattern.CategoryPrefix),
		Label:   l.Label,
	}
}
