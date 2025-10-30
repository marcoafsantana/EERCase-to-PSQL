package main

import (
	"eercase/dto"
	"eercase/printer"
	"eercase/sqlgen"
	"encoding/json"
	"log"
	"os"
)

func main() {
	// Lê o arquivo JSON
	jsonData, err := os.ReadFile("project.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// Cria uma instância da estrutura Project
	var project dto.ProjectRelationsDTO

	// Faz o unmarshal do JSON para a estrutura
	err = json.Unmarshal(jsonData, &project)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do JSON: %v", err)
	}
	relations := project.ToProjectEntity()
	// Imprime os detalhes do projeto usando o módulo printer
	if err := printer.PrintProjectDetails(*relations); err != nil {
		log.Fatalf("Erro ao imprimir detalhes do projeto: %v", err)
	}

	svc := sqlgen.NewService()
	if err := svc.GenerateSQLToFile(project, "tbl_create.sql"); err != nil {
		log.Fatalf("Erro ao gerar arquivo SQL: %v", err)
	}
}
