package nodes

import (
	nodes2 "eercase/models/eercase/nodes"
)

// Element: extende Node
type ElementDTO struct {
	NodeDTO
	Name string `gorm:"type:varchar(255)" json:"name"`
}

func NewElementDTOFromEntity(l *nodes2.Element, prefix string) ElementDTO {
	return ElementDTO{
		NodeDTO: NewNodeDTOFromEntity(&l.Node, prefix),
		Name:    l.Name,
	}
}
