package nodes

import (
	"eercase/models/eercase/enum"
	"eercase/pattern"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

// Inheritance: node com label e disjointness
type Inheritance struct {
	Node
	Label        string                `gorm:"type:varchar(255)" json:"label"`
	Disjointness enum.DisjointnessType `gorm:"not null" json:"disjointness" validate:"required"`
}

func (l *Inheritance) Create(db *gorm.DB) (uint, error) {
	if l.ID != 0 {
		return l.ID, errors.New("ID must be 0")
	}

	if err := db.Create(l).Error; err != nil {
		return 0, err
	}
	return l.ID, nil
}

func (l *Inheritance) GetID() uint {
	return l.ID
}

func (l *Inheritance) SetID(id uint) {
	l.ID = id
}

func (l *Inheritance) GetErrcaseID() string {
	// Retorna o ID do link como uma string numérica
	return pattern.InheritancePrefix + strconv.FormatUint(uint64(l.ID), 10)
}
