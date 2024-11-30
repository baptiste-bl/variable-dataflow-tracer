package main

import (
	"fmt"
)

func main() {
	// Initialisation de l'analyse à partir de la ligne 1
	CrawlFromLine(line)
}

// CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées).
func CrawlFromLine(line int) {
	fmt.Printf("Analyzing line: %d\n", line)

	// Condition pour simuler une fin de fichier à la ligne 10
	if line > 10 {
		fmt.Println("End of file reached.")
		return
	}

	// Simuler un appel récursif à une fonction interne
	AnalyzeFunction(line + 1)

}

// AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
func AnalyzeFunction(line int) {
	fmt.Printf("Entering function at line: %d\n", line)

	// Simuler un appel récursif à CrawlFromLine
	CrawlFromLine(line)
}
