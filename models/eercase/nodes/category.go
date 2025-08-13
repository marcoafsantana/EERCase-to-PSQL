package nodes

import (
	"errors"

	"gorm.io/gorm"
)

// Category: node com label
type Category struct {
	Node
	Label string `gorm:"type:varchar(255)" json:"label"`
}

func (l *Category) Create(db *gorm.DB) (uint, error) {
	if l.ID != 0 {
		return l.ID, errors.New("ID must be 0")
	}

	if err := db.Create(l).Error; err != nil {
		return 0, err
	}
	return l.ID, nil
}

func (l *Category) GetID() uint {
	return l.ID
}

func (l *Category) SetID(id uint) {
	l.ID = id
}

// func (l *Category) GetErrcaseID() string {
// 	// Retorna o ID do link como uma string numérica
// 	return pattern.CategoryPrefix + strconv.FormatUint(uint64(l.ID), 10)
// }
