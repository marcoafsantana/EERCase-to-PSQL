package printer

import (
	"eercase/dto"
)

// PrintProjectDetails imprime todos os detalhes do projeto de forma organizada
func PrintProjectDetails(project dto.ProjectRelationsDTO) error {
	printHeader("Informações do Projeto EER")

	printBasicInfo(project)
	printStats(project)
	printLinksSummary(project)
	printEntitiesDetails(project)
	printRelationships(project)
	printAssociativeEntities(project)

	return nil
}
