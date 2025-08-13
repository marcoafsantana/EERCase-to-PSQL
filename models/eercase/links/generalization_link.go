package links

import (
	enum2 "eercase/models/eercase/enum"
)

// GeneralizationLink: extende Link com role, completeness e type
type GeneralizationLink struct {
	Link
	Role         string                 `gorm:"type:varchar(255)" json:"role"`
	Completeness enum2.CompletenessType `gorm:"not null" json:"completeness" validate:"required"`
	Type         enum2.SLGLLinkType     `gorm:"not null" json:"type" validate:"required"`
}

// func (l *GeneralizationLink) GetErrcaseID() string {
// 	// Retorna o ID do link como uma string numérica
// 	return pattern.GeneralizationLinkPrefix + strconv.FormatUint(uint64(l.ID), 10)
// }
