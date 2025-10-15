package main

import (
	"eercase/models"
	"eercase/printer"
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
	var project models.Project

	// Faz o unmarshal do JSON para a estrutura
	err = json.Unmarshal(jsonData, &project)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do JSON: %v", err)
	}

	// Imprime os detalhes do projeto usando o módulo printer
	if err := printer.PrintProjectDetails(project); err != nil {
		log.Fatalf("Erro ao imprimir detalhes do projeto: %v", err)
	}
}
