package links

import (
	"eercase/pattern"
	"strconv"
)

// AttributeLink: extende Link sem campos adicionais
type AttributeLink struct {
	Link
}

func (l *AttributeLink) GetErrcaseID() string {
	// Retorna o ID do link como uma string numérica
	return pattern.AttributeLinkPrefix + strconv.FormatUint(uint64(l.ID), 10)
}
