package sqlgen

import (
	"eercase/dto"
	"eercase/dto/nodes"
	"eercase/models/eercase/enum"
	"fmt"
	"os"
	"strings"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Método no Service
func (s *Service) GenerateSQL(project dto.ProjectRelationsDTO) (string, error) {
	var sql strings.Builder

	sql.WriteString("-- Arquivo gerado automaticamente pelo EERCase-to-PSQL\n")
	sql.WriteString("-- Step 1: Criação das tabelas com seus atributos\n\n")

	// Step 1: cria definições das tabelas
	s.buildCreateTables(&sql, project)

	// Step 2: Adicionar chaves primárias para entidades fortes ou super-entidades
	sql.WriteString("-- Step 2: Adição das chaves primárias\n")
	s.buildPrimaryKeys(&sql, project)

	// Step 3: Para entidades fracas, adicionar colunas identificadoras do dono e montar PK composta
	sql.WriteString("-- Step 3: Identificação de entidades fracas (chaves compostas)\n")
	s.buildWeakEntityKeys(&sql, project)

	// Step 4: Para sub-entidades, adicionar colunas identificadoras do(s) super e montar PK composta
	sql.WriteString("-- Step 4: Identificação de sub-entidades (herança)\n")
	s.buildSubEntityKeys(&sql, project)

	// Escreve o arquivo SQL
	return sql.String(), nil
}

func (s *Service) GenerateSQLToFile(project dto.ProjectRelationsDTO, outputFile string) error {
	sql, err := s.GenerateSQL(project)
	if err != nil {
		return err
	}
	return os.WriteFile(outputFile, []byte(sql), 0o644)
}

// buildCreateTables escreve as instruções CREATE TABLE no builder
func (s *Service) buildCreateTables(sql *strings.Builder, project dto.ProjectRelationsDTO) {
	for _, entity := range project.Entities {
		sql.WriteString(fmt.Sprintf("-- Tabela %s\n", entity.Name))
		sql.WriteString(fmt.Sprintf("CREATE TABLE \"%s\" (\n", strings.ToLower(entity.Name)))

		// Encontra os atributos desta entidade
		var attributes []string

		// Adicionar atributos herdados (se a entidade for sub-entidade)
		inheritedAttrs := s.collectInheritedAttributes(entity, project)
		for _, attr := range inheritedAttrs {
			sqlType := s.getSQLType(attr)
			columnDef := fmt.Sprintf("    \"%s\" %s", strings.ToLower(attr.Name), sqlType)
			attributes = append(attributes, columnDef)
		}

		// Adicionar atributos próprios da entidade
		for _, attrLink := range project.AttributeLinks {
			if attrLink.SourceID != entity.ID {
				continue
			}
			for _, attr := range project.Attributes {
				if attr.ID != attrLink.TargetID {
					continue
				}
				// Determina o tipo SQL e constrói a definição da coluna
				sqlType := s.getSQLType(attr)
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
func (s *Service) buildPrimaryKeys(sql *strings.Builder, project dto.ProjectRelationsDTO) {
	for _, entity := range project.Entities {
		if !s.isSuperOrStrongEntity(entity, project) {
			continue
		}

		pkAttrs := s.collectPKAttrs(entity, project)
		if len(pkAttrs) == 0 {
			continue
		}

		sql.WriteString(fmt.Sprintf("ALTER TABLE \"%s\"\n", strings.ToLower(entity.Name)))
		sql.WriteString(fmt.Sprintf("ADD PRIMARY KEY (%s);\n\n", strings.Join(pkAttrs, ", ")))
	}
}

// buildWeakEntityKeys adiciona colunas de identificação do dono nas entidades fracas e cria PK compostas
func (s *Service) buildWeakEntityKeys(sql *strings.Builder, project dto.ProjectRelationsDTO) {
	for _, entity := range project.Entities {
		if !entity.IsWeak {
			continue
		}

		// encontrar relationships que identificam esta entidade
		// Para cada relationship com IsIdentifier true que conecta a entidade fraca com uma entidade forte
		ownerIdentifierAttrs := make(map[string]nodes.AttributeDTO)
		for _, rel := range project.Relationships {
			if !rel.IsIdentifier {
				continue
			}
			// coletar entidades ligadas a esta relationship
			var linkedEntityIDs []string
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
					// coletar atributos identificadores da entidade forte (próprios + herdados)
					allAttrs := s.collectAllAttributes(e, project)
					for _, attr := range allAttrs {
						if attr.Type == enum.IDENTIFIER {
							ownerIdentifierAttrs[attr.ID] = attr
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
			colType := s.getSQLType(attr)
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

// buildSubEntityKeys adiciona colunas identificadoras dos super-entidades para cada sub-entidade e cria PK composta
func (s *Service) buildSubEntityKeys(sql *strings.Builder, project dto.ProjectRelationsDTO) {
	for _, entity := range project.Entities {
		// verifica se é sub-entidade (existe um directInheritanceLink com TargetID == entity.ID)
		var inheritanceSourceID string
		for _, d := range project.DirectInheritanceLinks {
			if d.TargetID == entity.ID {
				inheritanceSourceID = d.SourceID
				break
			}
		}
		if inheritanceSourceID == "" {
			continue
		}

		// encontrar outras entidades ligadas ao mesmo inheritanceSourceID (possíveis super-entidades)
		superIdentifierAttrs := make(map[string]nodes.AttributeDTO)
		for _, d := range project.DirectInheritanceLinks {
			if d.SourceID != inheritanceSourceID {
				continue
			}
			if d.TargetID == entity.ID {
				continue
			}
			// d.TargetID é uma entidade relacionada à mesma herança — tratar como super-entidade
			for _, e := range project.Entities {
				if e.ID != d.TargetID {
					continue
				}
				if e.IsWeak {
					continue
				}
				// coletar atributos identificadores da entidade super
				for _, attrLink := range project.AttributeLinks {
					if attrLink.SourceID != e.ID {
						continue
					}
					for _, attr := range project.Attributes {
						if attr.ID != attrLink.TargetID {
							continue
						}
						if attr.Type == enum.IDENTIFIER {
							superIdentifierAttrs[attr.ID] = attr
						}
					}
				}
			}
		}

		if len(superIdentifierAttrs) == 0 {
			continue
		}

		// preparar ADD COLUMN para cada atributo identificador do super
		var addCols []string
		var pkParts []string
		for _, attr := range superIdentifierAttrs {
			colName := fmt.Sprintf("\"%s_super_pk\"", strings.ToLower(attr.Name))
			colType := s.getSQLType(attr)
			addCols = append(addCols, fmt.Sprintf("ADD COLUMN %s %s", colName, colType))
			pkParts = append(pkParts, colName)
		}

		// incluir identificadores próprios e discriminador da sub-entidade
		for _, attrLink := range project.AttributeLinks {
			if attrLink.SourceID != entity.ID {
				continue
			}
			for _, attr := range project.Attributes {
				if attr.ID != attrLink.TargetID {
					continue
				}
				if attr.Type == enum.IDENTIFIER {
					colName := fmt.Sprintf("\"%s\"", strings.ToLower(attr.Name))
					pkParts = append(pkParts, colName)
				}
				if attr.Type == enum.DISCRIMINATOR {
					colName := fmt.Sprintf("\"%s\"", strings.ToLower(attr.Name))
					pkParts = append(pkParts, colName)
				}
			}
		}

		// escrever ALTER TABLE ADD COLUMN ...
		sql.WriteString(fmt.Sprintf("ALTER TABLE \"%s\"\n", strings.ToLower(entity.Name)))
		if len(addCols) > 0 {
			sql.WriteString(strings.Join(addCols, ",\n"))
			sql.WriteString(";\n\n")
		}

		// escrever ALTER TABLE ADD PRIMARY KEY (...) se existirem partes de PK
		if len(pkParts) > 0 {
			sql.WriteString(fmt.Sprintf("ALTER TABLE \"%s\"\n", strings.ToLower(entity.Name)))
			sql.WriteString(fmt.Sprintf("ADD PRIMARY KEY (%s);\n\n", strings.Join(pkParts, ", ")))
		}
	}
}

// isSuperOrStrongEntity determina se a entidade é forte ou participa de hierarquia (super-entidade)
func (s *Service) isSuperOrStrongEntity(entity nodes.EntityDTO, project dto.ProjectRelationsDTO) bool {
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
func (s *Service) collectPKAttrs(entity nodes.EntityDTO, project dto.ProjectRelationsDTO) []string {
	var pkAttrs []string

	// Coletar todos os atributos (próprios + herdados)
	allAttrs := s.collectAllAttributes(entity, project)

	for _, attr := range allAttrs {
		if attr.Type == enum.IDENTIFIER {
			pkAttrs = append(pkAttrs, fmt.Sprintf("\"%s\"", strings.ToLower(attr.Name)))
		}
	}

	return pkAttrs
}

// collectInheritedAttributes coleta todos os atributos herdados de entidades pai através de hierarquia
func (s *Service) collectInheritedAttributes(entity nodes.EntityDTO, project dto.ProjectRelationsDTO) []nodes.AttributeDTO {
	var inheritedAttrs []nodes.AttributeDTO

	// Verificar se a entidade é uma especialização (sub-entidade)
	// Procurar por specialization_links onde source_id é a entidade atual
	var inheritanceNodeID string
	for _, specLink := range project.SpecializationLinks {
		if specLink.SourceID == entity.ID {
			// specLink.TargetID é o nó de herança
			inheritanceNodeID = specLink.TargetID
			break
		}
	}

	if inheritanceNodeID == "" {
		return inheritedAttrs
	}

	// Encontrar a entidade pai através do generalization_link
	// O generalization_link conecta a entidade pai ao nó de herança
	var parentEntityID string
	for _, genLink := range project.GeneralizationLinks {
		if genLink.TargetID == inheritanceNodeID {
			// genLink.SourceID é a entidade pai
			parentEntityID = genLink.SourceID
			break
		}
	}

	if parentEntityID == "" {
		return inheritedAttrs
	}

	// Coletar os atributos da entidade pai
	for _, attrLink := range project.AttributeLinks {
		if attrLink.SourceID != parentEntityID {
			continue
		}
		for _, attr := range project.Attributes {
			if attr.ID != attrLink.TargetID {
				continue
			}
			inheritedAttrs = append(inheritedAttrs, attr)
		}
	}

	return inheritedAttrs
}

// collectAllAttributes coleta todos os atributos de uma entidade (próprios + herdados)
func (s *Service) collectAllAttributes(entity nodes.EntityDTO, project dto.ProjectRelationsDTO) []nodes.AttributeDTO {
	var allAttrs []nodes.AttributeDTO

	// Primeiro, adicionar atributos herdados
	inheritedAttrs := s.collectInheritedAttributes(entity, project)
	allAttrs = append(allAttrs, inheritedAttrs...)

	// Depois, adicionar atributos próprios
	for _, attrLink := range project.AttributeLinks {
		if attrLink.SourceID != entity.ID {
			continue
		}
		for _, attr := range project.Attributes {
			if attr.ID != attrLink.TargetID {
				continue
			}
			allAttrs = append(allAttrs, attr)
		}
	}

	return allAttrs
}

// getSQLType converte o tipo de atributo EER para tipo SQL
func (s *Service) getSQLType(attr nodes.AttributeDTO) string {
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
