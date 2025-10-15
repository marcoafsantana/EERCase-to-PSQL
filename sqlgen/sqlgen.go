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

	// Gera as tabelas para as entidades
	for _, entity := range project.Entities {
		sql.WriteString(fmt.Sprintf("-- Tabela %s\n", entity.Name))
		sql.WriteString(fmt.Sprintf("CREATE TABLE \"%s\" (\n", strings.ToLower(entity.Name)))

		// Encontra os atributos desta entidade
		var attributes []string

		for _, attrLink := range project.AttributeLinks {
			if attrLink.SourceID == entity.ID {
				for _, attr := range project.Attributes {
					if attr.ID == attrLink.TargetID {
						// Determina o tipo SQL e constrói a definição da coluna
						sqlType := getSQLType(attr)
						columnDef := fmt.Sprintf("    \"%s\" %s", strings.ToLower(attr.Name), sqlType)
						attributes = append(attributes, columnDef)
					}
				}
			}
		}

		// Adiciona todas as definições de colunas
		sql.WriteString(strings.Join(attributes, ",\n"))
		sql.WriteString("\n);\n\n")
	}

	// Escreve o arquivo SQL
	return os.WriteFile(outputFile, []byte(sql.String()), 0644)
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
