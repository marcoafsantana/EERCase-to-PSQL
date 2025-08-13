package nodes

import (
	"errors"

	"gorm.io/gorm"
)

// Entity: extende Element
type Entity struct {
	Element
	IsWeak bool `gorm:"not null" json:"is_weak" validate:"required"`
}

func (l *Entity) Create(db *gorm.DB) (uint, error) {
	if l.ID != 0 {
		return l.ID, errors.New("ID must be 0")
	}
	if err := db.Create(l).Error; err != nil {
		return 0, err
	}
	return l.ID, nil
}

func (l *Entity) GetID() uint {
	return l.ID
}

func (l *Entity) SetID(id uint) {
	l.ID = id
}

// func (l *Entity) GetErrcaseID() string {
// 	// Retorna o ID do link como uma string numérica
// 	return pattern.EntityPrefix + strconv.FormatUint(uint64(l.ID), 10)
// }
