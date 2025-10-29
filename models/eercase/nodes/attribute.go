package nodes

import (
	enum2 "eercase/models/eercase/enum"
	"eercase/pattern"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

// Attribute: extende Element e utiliza enums do pacote enum
type Attribute struct {
	Element
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

func (l *Attribute) Create(db *gorm.DB) (uint, error) {
	if l.ID != 0 {
		return l.ID, errors.New("ID must be 0")
	}

	if err := db.Create(l).Error; err != nil {
		return 0, err
	}
	return l.ID, nil
}

func (l *Attribute) GetID() uint {
	return l.ID
}

func (l *Attribute) SetID(id uint) {
	l.ID = id
}

func (l *Attribute) GetErrcaseID() string {
	// Retorna o ID do link como uma string numérica
	return pattern.AttributePrefix + strconv.FormatUint(uint64(l.ID), 10)
}
