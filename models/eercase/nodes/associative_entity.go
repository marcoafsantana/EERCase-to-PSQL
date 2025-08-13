package nodes

import (
	"errors"

	"gorm.io/gorm"
)

// AssociativeEntity: extende Entity
type AssociativeEntity struct {
	Entity

	RelationshipID uint         `json:"relationship_id"`
	Relationship   Relationship `gorm:"foreignKey:RelationshipID,ProjectID;references:ID,ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

func (l *AssociativeEntity) Create(db *gorm.DB) (uint, error) {

	if l.ID != 0 {
		return l.ID, errors.New("ID must be 0")
	}

	if err := db.Create(l).Error; err != nil {
		return 0, err
	}

	return l.ID, nil
}

func (l *AssociativeEntity) GetID() uint {
	return l.ID
}

func (l *AssociativeEntity) SetID(id uint) {
	l.ID = id
}

// func (l *AssociativeEntity) GetErrcaseID() string {
// 	// Retorna o ID do link como uma string numérica
// 	return pattern.AssociativeEntityPrefix + strconv.FormatUint(uint64(l.ID), 10)
// }
