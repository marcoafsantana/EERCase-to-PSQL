package dto

import (
	"eercase/dto/links"
	"eercase/dto/nodes"
	"eercase/models"
)

// ProjectRelationsDTO agrupa o ID do projeto e todas as listas de links/nós
type ProjectRelationsDTO struct {
	SpecializationLinks    []links.SpecializationLinkDTO    `json:"specialization_links,omitempty"`
	GeneralizationLinks    []links.GeneralizationLinkDTO    `json:"generalization_links,omitempty"`
	DirectInheritanceLinks []links.DirectInheritanceLinkDTO `json:"direct_inheritance_links,omitempty"`
	AttributeLinks         []links.AttributeLinkDTO         `json:"attribute_links,omitempty"`
	RelationshipLinks      []links.RelationshipLinkDTO      `json:"relationship_links,omitempty"`
	Inheritances           []nodes.InheritanceDTO           `json:"inheritances,omitempty"`
	Categories             []nodes.CategoryDTO              `json:"categories,omitempty"`
	Entities               []nodes.EntityDTO                `json:"entities,omitempty"`
	Attributes             []nodes.AttributeDTO             `json:"attributes,omitempty"`
	Relationships          []nodes.RelationshipDTO          `json:"relationships,omitempty"`
	AssociativeEntities    []nodes.AssociativeEntityDTO     `json:"associative_entities,omitempty"`
}

func NewProjectRelationsDTOFromEntityAndPermission(project *models.Project) *ProjectRelationsDTO {
	dto := &ProjectRelationsDTO{}

	for _, link := range project.SpecializationLinks {
		dto.SpecializationLinks = append(dto.SpecializationLinks, links.NewSpecializationLinkDTOFromEntity(&link))
	}

	for _, link := range project.GeneralizationLinks {
		dto.GeneralizationLinks = append(dto.GeneralizationLinks, links.NewGeneralizationLinkDTOFromEntity(&link))
	}

	for _, link := range project.DirectInheritanceLinks {
		dto.DirectInheritanceLinks = append(dto.DirectInheritanceLinks, links.NewDirectInheritanceLinkDTOFromEntity(&link))
	}

	for _, link := range project.AttributeLinks {
		dto.AttributeLinks = append(dto.AttributeLinks, links.NewAttributeLinkDTOFromEntity(&link))
	}

	for _, link := range project.RelationshipLinks {
		dto.RelationshipLinks = append(dto.RelationshipLinks, links.NewRelationshipLinkDTOFromEntity(&link))
	}

	for _, inh := range project.Inheritances {
		dto.Inheritances = append(dto.Inheritances, nodes.NewInheritanceDTOFromEntity(&inh))
	}
	for _, cat := range project.Categories {
		dto.Categories = append(dto.Categories, nodes.NewCategoryDTOFromEntity(&cat))
	}
	for _, ent := range project.Entities {
		dto.Entities = append(dto.Entities, nodes.NewEntityDTOFromEntity(&ent))
	}
	for _, attr := range project.Attributes {
		dto.Attributes = append(dto.Attributes, nodes.NewAttributeDTOFromEntity(&attr))
	}
	for _, rel := range project.Relationships {
		dto.Relationships = append(dto.Relationships, nodes.NewRelationshipDTOFromEntity(&rel))
	}
	for _, assoc := range project.AssociativeEntities {
		dto.AssociativeEntities = append(dto.AssociativeEntities, nodes.NewAssociativeEntityDTOFromEntity(&assoc))
	}

	return dto
}
