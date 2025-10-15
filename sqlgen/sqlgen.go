package sqlgen

import (
	"eercase/models"
	"eercase/models/eercase/enum"
	"eercase/models/eercase/nodes"
	"fmt"
	"os"
	"strings"
)

// GenerateSQL gera o arquivo SQL para criar as tabelas baseadas no projeto EER
func GenerateSQL(project models.Project, outputFile string) error {
	var sql strings.Builder

	sql.WriteString("-- Arquivo gerado automaticamente pelo EERCase-to-PSQL\n")
	sql.WriteString("-- Step 1: Criação das tabelas com seus atributos\n\n")

	// Passo A: cria definições das tabelas
	buildCreateTables(&sql, project)

	// Step 2: Adicionar chaves primárias para entidades fortes ou super-entidades
	sql.WriteString("-- Step 2: Adição das chaves primárias\n")
	buildPrimaryKeys(&sql, project)

	// Step 3: Para entidades fracas, adicionar colunas identificadoras do dono e montar PK composta
	sql.WriteString("-- Step 3: Identificação de entidades fracas (chaves compostas)\n")
	buildWeakEntityKeys(&sql, project)

	// Escreve o arquivo SQL
	return os.WriteFile(outputFile, []byte(sql.String()), 0644)
}

// buildCreateTables escreve as instruções CREATE TABLE no builder
func buildCreateTables(sql *strings.Builder, project models.Project) {
	for _, entity := range project.Entities {
		sql.WriteString(fmt.Sprintf("-- Tabela %s\n", entity.Name))
		sql.WriteString(fmt.Sprintf("CREATE TABLE \"%s\" (\n", strings.ToLower(entity.Name)))

		// Encontra os atributos desta entidade
		var attributes []string
		for _, attrLink := range project.AttributeLinks {
			if attrLink.SourceID != entity.ID {
				continue
			}
			for _, attr := range project.Attributes {
				if attr.ID != attrLink.TargetID {
					continue
				}
				// Determina o tipo SQL e constrói a definição da coluna
				sqlType := getSQLType(attr)
				columnDef := fmt.Sprintf("    \"%s\" %s", strings.ToLower(attr.Name), sqlType)
				attributes = append(attributes, columnDef)
			}
		}

		// Adiciona todas as definições de colunas
		sql.WriteString(strings.Join(attributes, ",\n"))
		sql.WriteString("\n);\n\n")
	}
}

// buildPrimaryKeys escreve as instruções ALTER TABLE ... ADD PRIMARY KEY para cada entidade forte ou super-entidade
func buildPrimaryKeys(sql *strings.Builder, project models.Project) {
	for _, entity := range project.Entities {
		if !isSuperOrStrongEntity(entity, project) {
			continue
		}

		pkAttrs := collectPKAttrs(entity, project)
		if len(pkAttrs) == 0 {
			continue
		}

		sql.WriteString(fmt.Sprintf("ALTER TABLE \"%s\"\n", strings.ToLower(entity.Name)))
		sql.WriteString(fmt.Sprintf("ADD PRIMARY KEY (%s);\n\n", strings.Join(pkAttrs, ", ")))
	}
}

// buildWeakEntityKeys adiciona colunas de identificação do dono nas entidades fracas e cria PK compostas
func buildWeakEntityKeys(sql *strings.Builder, project models.Project) {
	for _, entity := range project.Entities {
		if !entity.IsWeak {
			continue
		}

		// encontrar relationships que identificam esta entidade
		// Para cada relationship com IsIdentifier true que conecta a entidade fraca com uma entidade forte
		ownerIdentifierAttrs := make(map[uint]nodes.Attribute) // map[attr.ID]attr
		for _, rel := range project.Relationships {
			if !rel.IsIdentifier {
				continue
			}
			// coletar entidades ligadas a esta relationship
			var linkedEntityIDs []uint
			for _, rl := range project.RelationshipLinks {
				if rl.TargetID == rel.ID {
					linkedEntityIDs = append(linkedEntityIDs, rl.SourceID)
				}
			}

			// se a entidade fraca participa desta relationship
			participates := false
			for _, id := range linkedEntityIDs {
				if id == entity.ID {
					participates = true
					break
				}
			}
			if !participates {
				continue
			}

			// para cada outra entidade ligada que seja forte, coletar seus atributos identificadores
			for _, id := range linkedEntityIDs {
				if id == entity.ID {
					continue
				}
				// encontrar a entidade
				for _, e := range project.Entities {
					if e.ID != id {
						continue
					}
					if e.IsWeak {
						continue
					}
					// coletar atributos identificadores da entidade forte
					for _, attrLink := range project.AttributeLinks {
						if attrLink.SourceID != e.ID {
							continue
						}
						for _, attr := range project.Attributes {
							if attr.ID != attrLink.TargetID {
								continue
							}
							if attr.Type == enum.IDENTIFIER {
								ownerIdentifierAttrs[attr.ID] = attr
							}
						}
					}
				}
			}
		}

		// se não encontrou atributos identificadores no dono, pula
		if len(ownerIdentifierAttrs) == 0 {
			continue
		}

		// prepara instrução ADD COLUMN para cada atributo identificador do dono
		var addCols []string
		var pkParts []string
		// adicionar os atributos do dono como "nome_pk"
		for _, attr := range ownerIdentifierAttrs {
			colName := fmt.Sprintf("\"%s_pk\"", strings.ToLower(attr.Name))
			colType := getSQLType(attr)
			addCols = append(addCols, fmt.Sprintf("ADD COLUMN %s %s", colName, colType))
			pkParts = append(pkParts, colName)
		}

		// coletar atributos identificadores próprios da entidade fraca
		for _, attrLink := range project.AttributeLinks {
			if attrLink.SourceID != entity.ID {
				continue
			}
			for _, attr := range project.Attributes {
				if attr.ID != attrLink.TargetID {
					continue
				}
				if attr.Type == enum.IDENTIFIER {
					// nome existente na tabela
					colName := fmt.Sprintf("\"%s\"", strings.ToLower(attr.Name))
					pkParts = append(pkParts, colName)
				}
				// se existir um atributo discriminador, adiciona no pk também
				if attr.Type == enum.DISCRIMINATOR {
					colName := fmt.Sprintf("\"%s\"", strings.ToLower(attr.Name))
					pkParts = append(pkParts, colName)
				}
			}
		}

		// escrever ALTER TABLE ADD COLUMN ... (várias colunas)
		sql.WriteString(fmt.Sprintf("ALTER TABLE \"%s\"\n", strings.ToLower(entity.Name)))
		if len(addCols) > 0 {
			sql.WriteString(strings.Join(addCols, ",\n"))
			sql.WriteString(";\n\n")
		}

		// escrever ALTER TABLE ADD PRIMARY KEY (...)
		if len(pkParts) > 0 {
			sql.WriteString(fmt.Sprintf("ALTER TABLE \"%s\"\n", strings.ToLower(entity.Name)))
			sql.WriteString(fmt.Sprintf("ADD PRIMARY KEY (%s);\n\n", strings.Join(pkParts, ", ")))
		}
	}
}

// isSuperOrStrongEntity determina se a entidade é forte ou participa de hierarquia (super-entidade)
func isSuperOrStrongEntity(entity nodes.Entity, project models.Project) bool {
	if !entity.IsWeak {
		return true
	}
	for _, d := range project.DirectInheritanceLinks {
		if d.TargetID == entity.ID || d.SourceID == entity.ID {
			return true
		}
	}
	for _, s := range project.SpecializationLinks {
		if s.SourceID == entity.ID || s.TargetID == entity.ID {
			return true
		}
	}
	for _, g := range project.GeneralizationLinks {
		if g.SourceID == entity.ID || g.TargetID == entity.ID {
			return true
		}
	}
	return false
}

// collectPKAttrs retorna a lista de nomes de colunas (entre aspas) que são do tipo IDENTIFIER para a entidade
func collectPKAttrs(entity nodes.Entity, project models.Project) []string {
	var pkAttrs []string
	for _, attrLink := range project.AttributeLinks {
		if attrLink.SourceID != entity.ID {
			continue
		}
		for _, attr := range project.Attributes {
			if attr.ID != attrLink.TargetID {
				continue
			}
			if attr.Type == enum.IDENTIFIER {
				pkAttrs = append(pkAttrs, fmt.Sprintf("\"%s\"", strings.ToLower(attr.Name)))
			}
		}
	}
	return pkAttrs
}

// getSQLType converte o tipo de atributo EER para tipo SQL
func getSQLType(attr nodes.Attribute) string {
	switch attr.DataType {
	case enum.STRING:
		if attr.Size > 0 {
			return fmt.Sprintf("VARCHAR(%d)", int(attr.Size))
		}
		return "TEXT"
	case enum.INTEGER:
		return "INTEGER"
	case enum.FLOAT:
		return "DECIMAL"
	case enum.BOOLEAN:
		return "BOOLEAN"
	case enum.TIMESTAMP:
		return "TIMESTAMP"
	case enum.CLOB:
		return "TEXT"
	case enum.BLOB:
		return "BYTEA"
	default:
		return "TEXT"
	}
}
