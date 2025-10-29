package links

import (
	"eercase/pattern"
	"strconv"
)

// DirectInheritanceLink: extende Link com role
type DirectInheritanceLink struct {
	Link
	Role string `gorm:"type:varchar(255)" json:"role"`
}

func (l *DirectInheritanceLink) GetErrcaseID() string {
	// Retorna o ID do link como uma string numérica
	return pattern.DirectInheritanceLinkPrefix + strconv.FormatUint(uint64(l.ID), 10)
}
