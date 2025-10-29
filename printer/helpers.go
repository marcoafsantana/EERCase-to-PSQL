package printer

import (
	"eercase/models"
	"eercase/models/eercase/nodes"
	"fmt"
	"strings"
)

func printBasicInfo(project models.Project) {
	fmt.Printf("📋 Detalhes Básicos:\n")
	fmt.Printf("   • ID: %s\n", project.ID)
	fmt.Printf("   • Título: %s\n", project.Title)
	fmt.Printf("   • Proprietário: %s\n\n", project.Owner)
}

func printStats(project models.Project) {
	printHeader("Estatísticas do Modelo")
	fmt.Printf("📊 Elementos do Modelo:\n")
	fmt.Printf("   • Entidades: %d\n", len(project.Entities))
	fmt.Printf("   • Relacionamentos: %d\n", len(project.Relationships))
	fmt.Printf("   • Atributos: %d\n", len(project.Attributes))
	fmt.Printf("   • Entidades Associativas: %d\n", len(project.AssociativeEntities))
}

func printLinksSummary(project models.Project) {
	fmt.Printf("\n🔗 Links e Conexões:\n")
	fmt.Printf("   • Links de Especialização: %d\n", len(project.SpecializationLinks))
	fmt.Printf("   • Links de Generalização: %d\n", len(project.GeneralizationLinks))
	fmt.Printf("   • Links de Herança Direta: %d\n", len(project.DirectInheritanceLinks))
	fmt.Printf("   • Links de Relacionamento: %d\n", len(project.RelationshipLinks))
	fmt.Printf("   • Links de Atributos: %d\n", len(project.AttributeLinks))
}

func printEntitiesDetails(project models.Project) {
	if len(project.Entities) == 0 {
		return
	}

	printHeader("Detalhes das Entidades")
	for _, entity := range project.Entities {
		printEntityDetails(entity, project)
	}
}

func printEntityDetails(entity nodes.Entity, project models.Project) {
	fmt.Printf("\n🗃️  Entidade: %s\n", entity.Name)
	if entity.IsWeak {
		fmt.Printf("   ⚠️  Status: Entidade Fraca\n")
	}

	// Atributos
	fmt.Printf("   📝 Atributos:\n")
	hasAttributes := false
	for _, attrLink := range project.AttributeLinks {
		if attrLink.SourceID == entity.GetErrcaseID() {
			for _, attr := range project.Attributes {
				if attr.GetErrcaseID() == attrLink.TargetID {
					hasAttributes = true
					printAttributeDetails(attr)
				}
			}
		}
	}
	if !hasAttributes {
		fmt.Println("      ❌ Nenhum atributo definido")
	}

	// Relacionamentos da entidade
	fmt.Printf("\n   🔗 Relacionamentos:\n")
	hasRelationships := false
	for _, relLink := range project.RelationshipLinks {
		if relLink.SourceID == entity.GetErrcaseID() {
			for _, rel := range project.Relationships {
				if rel.GetErrcaseID() == relLink.TargetID {
					hasRelationships = true
					// Determinar cardinalidade
					cardinalidade := "1:1"
					if relLink.Cardinality == 1 {
						cardinalidade = "1:N"
					} else if relLink.Cardinality == 2 {
						cardinalidade = "N:M"
					}

					// Determinar participação
					participacao := "Parcial"
					if relLink.Participation == 1 {
						participacao = "Total"
					}

					fmt.Printf("      • %s", rel.Name)
					fmt.Printf(" (%s, %s)", cardinalidade, participacao)
					if relLink.IsIdentifier {
						fmt.Printf(" [Identificador]")
					}
					fmt.Println()
				}
			}
		}
	}
	if !hasRelationships {
		fmt.Println("      ❌ Nenhum relacionamento definido")
	}

	// Heranças
	fmt.Printf("\n   👨‍👦 Heranças:\n")
	hasInheritance := false
	for _, specLink := range project.SpecializationLinks {
		if specLink.SourceID == entity.GetErrcaseID() {
			hasInheritance = true
			fmt.Println("      • Especialização de outra entidade")
		}
	}
	for _, genLink := range project.GeneralizationLinks {
		if genLink.SourceID == entity.GetErrcaseID() {
			hasInheritance = true
			completeness := "Parcial"
			if genLink.Completeness == 1 {
				completeness = "Total"
			}
			fmt.Printf("      • Generalização (%s)\n", completeness)
		}
	}
	if !hasInheritance {
		fmt.Println("      ❌ Não participa de hierarquia")
	}

	fmt.Println() // Linha em branco entre entidades
}

func printAttributeDetails(attr nodes.Attribute) {
	// Determinar o tipo do atributo
	attrType := "Comum"
	switch attr.Type {
	case 1:
		attrType = "Derivado"
	case 2:
		attrType = "Multivalorado"
	case 3:
		attrType = "Chave"
	case 4:
		attrType = "Discriminador"
	}

	fmt.Printf("      • %s", attr.Name)
	if attrType != "Comum" {
		fmt.Printf(" (%s)", attrType)
	}
	if attr.IsUnique {
		fmt.Printf(" [Único]")
	}
	fmt.Println()

	fmt.Printf("        Tipo EER: %s\n", attrType)
	fmt.Printf("        Tipo de Dado: ")
	switch attr.DataType {
	case 0:
		fmt.Println("STRING")
	case 1:
		fmt.Println("BOOLEAN")
	case 2:
		fmt.Println("TIMESTAMP")
	case 3:
		fmt.Println("FLOAT")
	case 4:
		fmt.Println("INTEGER")
	case 5:
		fmt.Println("CLOB")
	case 6:
		fmt.Println("BLOB")
	default:
		fmt.Println("Desconhecido")
	}

	if attr.Size > 0 {
		fmt.Printf("        Tamanho: %.0f\n", attr.Size)
	}

	// Obrigatoriedade e unicidade
	if !attr.IsNull {
		fmt.Println("        Obrigatório: Sim")
	} else {
		fmt.Println("        Obrigatório: Não")
	}
	if attr.IsUnique {
		fmt.Println("        Único: Sim")
	} else {
		fmt.Println("        Único: Não")
	}

	// Cardinalidade
	cardLabel := "1"
	switch attr.Cardinality {
	case 0:
		cardLabel = "1"
	case 1:
		cardLabel = "N"
	default:
		cardLabel = fmt.Sprintf("%d", attr.Cardinality)
	}
	fmt.Printf("        Cardinalidade: %s\n", cardLabel)
	if attr.Cardinality != 0 {
		fmt.Println("        Observação: Atributo multivalorado")
	}

	if attr.DefaultValue != "" {
		fmt.Printf("        Valor Padrão: %s\n", attr.DefaultValue)
	}
	if attr.Check != "" {
		fmt.Printf("        Restrição: %s\n", attr.Check)
	}
	if attr.Comment != "" {
		fmt.Printf("        Comentário: %s\n", attr.Comment)
	}
}

func printRelationships(project models.Project) {
	if len(project.Relationships) == 0 {
		return
	}
	printHeader("Relacionamentos")
	for _, rel := range project.Relationships {
		fmt.Printf("🔗 %s\n", rel.Name)
	}
}

func printAssociativeEntities(project models.Project) {
	if len(project.AssociativeEntities) == 0 {
		return
	}
	printHeader("Entidades Associativas")
	for _, assoc := range project.AssociativeEntities {
		fmt.Printf("📦 %s\n", assoc.Name)
		if assoc.IsWeak {
			fmt.Printf("   ⚠️ Entidade Associativa Fraca\n")
		}
	}
}

// printHeader imprime um cabeçalho formatado
func printHeader(title string) {
	fmt.Printf("\n%s\n", title)
	fmt.Println(strings.Repeat("=", len(title)))
}
