package links

import (
	enum2 "eercase/models/eercase/enum"
	links2 "eercase/models/eercase/links"
	"eercase/pattern"
	"strconv"
	"strings"
)

// GeneralizationLink: extende Link com role, completeness e type
type GeneralizationLinkDTO struct {
	LinkDTO
	Role         string                 `gorm:"type:varchar(255)" json:"role"`
	Completeness enum2.CompletenessType `gorm:"not null" json:"completeness" validate:"required"`
	Type         enum2.SLGLLinkType     `gorm:"not null" json:"type" validate:"required"`
}

func (l *GeneralizationLinkDTO) ToEntity(nodeMap map[string]string) links2.GeneralizationLink {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.GeneralizationLinkPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.GeneralizationLinkPrefix)
	} else {
		rawID = l.ID
	}

	uid64, err := strconv.ParseUint(rawID, 10, 64)

	var id uint
	if err == nil {
		id = uint(uid64)
	} else {
		id = 0 // cria novo se a string não for numérica
	}

	return links2.GeneralizationLink{
		Link: links2.Link{
			ID:           id,
			ProjectID:    l.ProjectID,
			SourceID:     resolveID(l.SourceID, nodeMap),
			TargetID:     resolveID(l.TargetID, nodeMap),
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role:         l.Role,
		Type:         l.Type,
		Completeness: l.Completeness,
	}
}

func NewGeneralizationLinkDTOFromEntity(l *links2.GeneralizationLink) GeneralizationLinkDTO {
	return GeneralizationLinkDTO{
		LinkDTO: LinkDTO{
			ID:           pattern.GeneralizationLinkPrefix + strconv.FormatUint(uint64(l.ID), 10),
			ProjectID:    l.ProjectID,
			SourceID:     l.SourceID,
			TargetID:     l.TargetID,
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role:         l.Role,
		Type:         l.Type,
		Completeness: l.Completeness,
	}
}
