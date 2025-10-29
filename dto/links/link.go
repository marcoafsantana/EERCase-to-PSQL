package links

import (
	links2 "eercase/models/eercase/links"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Link: Todo link deve estar atrelado a um projeto
type LinkDTO struct {
	ID           string    `json:"id"`
	ProjectID    uuid.UUID `json:"project_id" validate:"required,uuid4"`
	SourceID     string    `validate:"required" json:"source_id"`
	TargetID     string    `validate:"required" json:"target_id"`
	SourceHandle string    `gorm:"not null" json:"source_handle" validate:"required"`
	TargetHandle string    `gorm:"not null" json:"target_handle" validate:"required"`
}

func (l *LinkDTO) ToEntity(nodeMap map[string]string, prefix string) links2.Link {
	// --- trata o ID vindo do front ------------------------------------
	// Se vier com o prefixo (ex.: "REL_123"), remove-o antes de converter.
	var rawID string
	if strings.HasPrefix(l.ID, prefix) {
		rawID = strings.TrimPrefix(l.ID, prefix)
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

	return links2.Link{
		ID:           id,
		ProjectID:    l.ProjectID,
		SourceID:     resolveID(l.SourceID, nodeMap),
		TargetID:     resolveID(l.TargetID, nodeMap),
		SourceHandle: l.SourceHandle,
		TargetHandle: l.TargetHandle,
	}
}
