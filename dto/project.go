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

func (dto *ProjectRelationsDTO) ToProjectEntity() *models.Project {
	project := &models.Project{}
	nodeMap := make(map[string]string)

	for _, link := range dto.SpecializationLinks {
		project.SpecializationLinks = append(project.SpecializationLinks, link.ToEntity(nodeMap))
	}

	for _, link := range dto.GeneralizationLinks {
		project.GeneralizationLinks = append(project.GeneralizationLinks, link.ToEntity(nodeMap))
	}

	for _, link := range dto.DirectInheritanceLinks {
		project.DirectInheritanceLinks = append(project.DirectInheritanceLinks, link.ToEntity(nodeMap))
	}

	for _, link := range dto.AttributeLinks {
		project.AttributeLinks = append(project.AttributeLinks, link.ToEntity(nodeMap))
	}

	for _, link := range dto.RelationshipLinks {
		project.RelationshipLinks = append(project.RelationshipLinks, link.ToEntity(nodeMap))
	}

	for _, inh := range dto.Inheritances {
		project.Inheritances = append(project.Inheritances, inh.ToEntity())
	}
	for _, cat := range dto.Categories {
		project.Categories = append(project.Categories, cat.ToEntity())
	}
	for _, ent := range dto.Entities {
		project.Entities = append(project.Entities, ent.ToEntity())
	}
	for _, attr := range dto.Attributes {
		project.Attributes = append(project.Attributes, attr.ToEntity())
	}
	for _, rel := range dto.Relationships {
		project.Relationships = append(project.Relationships, rel.ToEntity())
	}
	for _, assoc := range dto.AssociativeEntities {
		project.AssociativeEntities = append(project.AssociativeEntities, assoc.ToEntity(nodeMap))
	}

	return project
}
