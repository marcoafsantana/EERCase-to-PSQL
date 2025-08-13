package main

import (
	"eercase/models"
	"encoding/json"
	"fmt"
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

	// Converte a estrutura de volta para JSON com indentação para melhor visualização
	prettyJSON, err := json.MarshalIndent(project, "", "    ")
	if err != nil {
		log.Fatalf("Erro ao converter estrutura para JSON: %v", err)
	}

	// Imprime o resultado
	fmt.Println("Projeto carregado com sucesso:")
	fmt.Println(string(prettyJSON))

	// Imprime algumas informações específicas para validação
	fmt.Printf("\nInformações do Projeto:\n")
	fmt.Printf("ID: %s\n", project.ID)
	fmt.Printf("Título: %s\n", project.Title)
	fmt.Printf("Owner: %s\n", project.Owner)
	fmt.Printf("Número de Categories: %d\n", len(project.Categories))
	fmt.Printf("Número de Entities: %d\n", len(project.Entities))
	fmt.Printf("Número de Attributes: %d\n", len(project.Attributes))
	fmt.Printf("Número de Relationships: %d\n", len(project.Relationships))
	fmt.Printf("Número de Links de Especialização: %d\n", len(project.SpecializationLinks))
}
