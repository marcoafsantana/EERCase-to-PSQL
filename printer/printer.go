package printer

import (
	"eercase/models"
	"fmt"
	"strings"
)

// PrintProjectDetails imprime todos os detalhes do projeto de forma organizada
func PrintProjectDetails(project models.Project) error {
	printHeader("Informações do Projeto EER")

	// Informações básicas
	fmt.Printf("📋 Detalhes Básicos:\n")
	fmt.Printf("   • ID: %s\n", project.ID)
	fmt.Printf("   • Título: %s\n", project.Title)
	fmt.Printf("   • Proprietário: %s\n\n", project.Owner)

	// Estatísticas do projeto
	printHeader("Estatísticas do Modelo")
	fmt.Printf("📊 Elementos do Modelo:\n")
	fmt.Printf("   • Entidades: %d\n", len(project.Entities))
	fmt.Printf("   • Relacionamentos: %d\n", len(project.Relationships))
	fmt.Printf("   • Atributos: %d\n", len(project.Attributes))
	fmt.Printf("   • Entidades Associativas: %d\n", len(project.AssociativeEntities))

	fmt.Printf("\n🔗 Links e Conexões:\n")
	fmt.Printf("   • Links de Especialização: %d\n", len(project.SpecializationLinks))
	fmt.Printf("   • Links de Generalização: %d\n", len(project.GeneralizationLinks))
	fmt.Printf("   • Links de Herança Direta: %d\n", len(project.DirectInheritanceLinks))
	fmt.Printf("   • Links de Relacionamento: %d\n", len(project.RelationshipLinks))
	fmt.Printf("   • Links de Atributos: %d\n", len(project.AttributeLinks))

	// Detalhes das entidades
	if len(project.Entities) > 0 {
		printHeader("Detalhes das Entidades")
		for _, entity := range project.Entities {
			fmt.Printf("\n🗃️  Entidade: %s\n", entity.Name)
			if entity.IsWeak {
				fmt.Printf("   ⚠️  Status: Entidade Fraca\n")
			}

			// Atributos da entidade
			fmt.Printf("   📝 Atributos:\n")
			hasAttributes := false
			for _, attrLink := range project.AttributeLinks {
				if attrLink.SourceID == entity.ID {
					for _, attr := range project.Attributes {
						if attr.ID == attrLink.TargetID {
							hasAttributes = true
							// Determinar o tipo do atributo
							attrType := "Regular"
							switch attr.Type {
							case 1:
								attrType = "Derivado"
							case 2:
								attrType = "Multivalorado"
							case 3:
								attrType = "Chave"
							}

							fmt.Printf("      • %s", attr.Name)
							if attrType != "Regular" {
								fmt.Printf(" (%s)", attrType)
							}
							if attr.IsUnique {
								fmt.Printf(" [Único]")
							}
							fmt.Println()

							// Detalhes adicionais do atributo
							if attr.DataType != 0 {
								fmt.Printf("        Tipo de Dado: ")
								switch attr.DataType {
								case 1:
									fmt.Println("Número")
								case 2:
									fmt.Println("Data")
								case 3:
									fmt.Println("Booleano")
								default:
									fmt.Println("Texto")
								}
							}
							if attr.Size > 0 {
								fmt.Printf("        Tamanho: %.0f\n", attr.Size)
							}
							if !attr.IsNull {
								fmt.Println("        Obrigatório: Sim")
							}
							if attr.DefaultValue != "" {
								fmt.Printf("        Valor Padrão: %s\n", attr.DefaultValue)
							}
							if attr.Check != "" {
								fmt.Printf("        Restrição: %s\n", attr.Check)
							}
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
				if relLink.SourceID == entity.ID {
					for _, rel := range project.Relationships {
						if rel.ID == relLink.TargetID {
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

			// Se houver herança
			fmt.Printf("\n   👨‍👦 Heranças:\n")
			hasInheritance := false
			// Verificar especializações
			for _, specLink := range project.SpecializationLinks {
				if specLink.SourceID == entity.ID {
					hasInheritance = true
					fmt.Println("      • Especialização de outra entidade")
				}
			}
			// Verificar generalizações
			for _, genLink := range project.GeneralizationLinks {
				if genLink.SourceID == entity.ID {
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
	}

	// Detalhes dos relacionamentos
	if len(project.Relationships) > 0 {
		printHeader("Relacionamentos")
		for _, rel := range project.Relationships {
			fmt.Printf("🔗 %s\n", rel.Name)
		}
	}

	// Detalhes das entidades associativas
	if len(project.AssociativeEntities) > 0 {
		printHeader("Entidades Associativas")
		for _, assoc := range project.AssociativeEntities {
			fmt.Printf("📦 %s\n", assoc.Name)
			if assoc.IsWeak {
				fmt.Printf("   ⚠️ Entidade Associativa Fraca\n")
			}
		}
	}

	return nil
}

// printHeader imprime um cabeçalho formatado
func printHeader(title string) {
	fmt.Printf("\n%s\n", title)
	fmt.Println(strings.Repeat("=", len(title)))
}
