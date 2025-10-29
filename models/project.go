package models

import (
	"eercase/models/eercase/links"
	"eercase/models/eercase/nodes"

	"github.com/google/uuid"
)

// Project: contém os relacionamentos com outros elementos
type Project struct {
	ID    uuid.UUID `gorm:"type:char(36);primaryKey" json:"id" validate:"required,uuid4"`
	Title string    `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Owner string    `gorm:"type:varchar(255);not null;index" json:"owner" validate:"required,email"`

	SpecializationLinks    []links.SpecializationLink    `gorm:"foreignKey:ProjectID" json:"specialization_links,omitempty"`
	GeneralizationLinks    []links.GeneralizationLink    `gorm:"foreignKey:ProjectID" json:"generalization_links,omitempty"`
	DirectInheritanceLinks []links.DirectInheritanceLink `gorm:"foreignKey:ProjectID" json:"direct_inheritance_links,omitempty"`
	RelationshipLinks      []links.RelationshipLink      `gorm:"foreignKey:ProjectID" json:"relationship_links,omitempty"`
	AttributeLinks         []links.AttributeLink         `gorm:"foreignKey:ProjectID" json:"attribute_links,omitempty"`

	Inheritances        []nodes.Inheritance       `gorm:"foreignKey:ProjectID" json:"inheritances,omitempty"`
	Categories          []nodes.Category          `gorm:"foreignKey:ProjectID" json:"categories,omitempty"`
	Entities            []nodes.Entity            `gorm:"foreignKey:ProjectID" json:"entities,omitempty"`
	Attributes          []nodes.Attribute         `gorm:"foreignKey:ProjectID" json:"attributes,omitempty"`
	Relationships       []nodes.Relationship      `gorm:"foreignKey:ProjectID" json:"relationships,omitempty"`
	AssociativeEntities []nodes.AssociativeEntity `gorm:"foreignKey:ProjectID" json:"associative_entities,omitempty"`
}
