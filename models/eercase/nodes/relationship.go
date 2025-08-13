package nodes

import (
	"errors"

	"gorm.io/gorm"
)

// Relationship: extende Element
type Relationship struct {
	Element
	IsIdentifier bool `gorm:"not null" json:"is_identifier" validate:"required"`
}

func (l *Relationship) Create(db *gorm.DB) (uint, error) {
	if l.ID != 0 {
		return l.ID, errors.New("ID must be 0")
	}

	if err := db.Create(l).Error; err != nil {
		return 0, err
	}
	return l.ID, nil
}

func (l *Relationship) GetID() uint {
	return l.ID
}

func (l *Relationship) SetID(id uint) {
	l.ID = id
}

// func (l *Relationship) GetErrcaseID() string {
// 	// Retorna o ID do link como uma string numérica
// 	return pattern.RelationshipPrefix + strconv.FormatUint(uint64(l.ID), 10)
// }
