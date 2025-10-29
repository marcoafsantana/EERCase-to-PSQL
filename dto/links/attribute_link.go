package links

import (
	links2 "eercase/models/eercase/links"
	"eercase/pattern"
	"strconv"
	"strings"
)

// AttributeLink: extende Link sem campos adicionais
type AttributeLinkDTO struct {
	LinkDTO
}

func (l *AttributeLinkDTO) ToEntity(nodeMap map[string]string) links2.AttributeLink {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, pattern.AttributeLinkPrefix) {
		rawID = strings.TrimPrefix(l.ID, pattern.AttributeLinkPrefix)
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

	return links2.AttributeLink{
		Link: links2.Link{
			ID:           id,
			ProjectID:    l.ProjectID,
			SourceID:     resolveID(l.SourceID, nodeMap),
			TargetID:     resolveID(l.TargetID, nodeMap),
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
	}
}

func NewAttributeLinkDTOFromEntity(l *links2.AttributeLink) AttributeLinkDTO {
	return AttributeLinkDTO{
		LinkDTO: LinkDTO{
			ID:           pattern.AttributeLinkPrefix + strconv.FormatUint(uint64(l.ID), 10),
			ProjectID:    l.ProjectID,
			SourceID:     l.SourceID,
			TargetID:     l.TargetID,
			SourceHandle: l.SourceHandle,
			TargetHandle: l.TargetHandle,
		},
	}
}
