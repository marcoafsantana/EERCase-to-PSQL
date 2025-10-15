package printer

import (
	"eercase/models"
)

// PrintProjectDetails imprime todos os detalhes do projeto de forma organizada
func PrintProjectDetails(project models.Project) error {
	printHeader("Informações do Projeto EER")

	printBasicInfo(project)
	printStats(project)
	printLinksSummary(project)
	printEntitiesDetails(project)
	printRelationships(project)
	printAssociativeEntities(project)

	return nil
}
