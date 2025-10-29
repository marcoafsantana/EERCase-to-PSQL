package links

import (
	"eercase/models/eercase/enum"
	links2 "eercase/models/eercase/links"
	"eercase/pattern"
	"strconv"
	"strings"
)

// SpecializationLink: extende Link com role e type
type SpecializationLinkDTO struct {
	LinkDTO
	Role string            `gorm:"type:varchar(255)" json:"role"`
	Type enum.SLGLLinkType `gorm:"not null" json:"type" validate:"required"`
}

func (l *SpecializationLinkDTO) ToEntity(nodeMap map[string]string) links2.SpecializationLink {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.SpecializationLinkPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.SpecializationLinkPrefix)
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

	return links2.SpecializationLink{
		Link: links2.Link{
			ID:           id,
			ProjectID:    l.ProjectID,
			SourceID:     resolveID(l.SourceID, nodeMap),
			TargetID:     resolveID(l.TargetID, nodeMap),
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role: l.Role,
		Type: l.Type,
	}
}

func NewSpecializationLinkDTOFromEntity(l *links2.SpecializationLink) SpecializationLinkDTO {
	return SpecializationLinkDTO{
		LinkDTO: LinkDTO{
			ID:           pattern.SpecializationLinkPrefix + strconv.FormatUint(uint64(l.ID), 10),
			ProjectID:    l.ProjectID,
			SourceID:     l.SourceID,
			TargetID:     l.TargetID,
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role: l.Role,
		Type: l.Type,
	}
}
