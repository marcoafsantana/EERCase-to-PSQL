package links

import (
	links2 "eercase/models/eercase/links"
	"eercase/pattern"
	"strconv"
	"strings"
)

// DirectInheritanceLink: extende Link com role
type DirectInheritanceLinkDTO struct {
	LinkDTO
	Role string `gorm:"type:varchar(255)" json:"role"`
}

func (l *DirectInheritanceLinkDTO) ToEntity(nodeMap map[string]string) links2.DirectInheritanceLink {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.DirectInheritanceLinkPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.DirectInheritanceLinkPrefix)
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

	return links2.DirectInheritanceLink{
		Link: links2.Link{
			ID:           id,
			ProjectID:    l.ProjectID,
			SourceID:     resolveID(l.SourceID, nodeMap),
			TargetID:     resolveID(l.TargetID, nodeMap),
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role: l.Role,
	}
}

func NewDirectInheritanceLinkDTOFromEntity(l *links2.DirectInheritanceLink) DirectInheritanceLinkDTO {
	return DirectInheritanceLinkDTO{
		LinkDTO: LinkDTO{
			ID:           pattern.DirectInheritanceLinkPrefix + strconv.FormatUint(uint64(l.ID), 10),
			ProjectID:    l.ProjectID,
			SourceID:     l.SourceID,
			TargetID:     l.TargetID,
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
		Role: l.Role,
	}
}
