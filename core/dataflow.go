package core

import (
	"dataflow/crawler"
	"dataflow/logger"
	"dataflow/models"
	"dataflow/services/dataFlowService"
	"dataflow/services/languageService"
	"dataflow/services/nodeService"
	"fmt"
	"log"
	"os"
)

// -----------------------------------------------------------------------------
// RunDataflowAnalysis - Runs data flow analysis based on the provided configuration
// -----------------------------------------------------------------------------
//
// Parameters:
//   - config (models.Config): Configuration settings for the data flow analysis.
//
// Returns:
//   - ([]models.DataFlow): A slice of DataFlow models representing the result of the analysis.
//   - (error): An error object if an error occurred during the analysis.
//
// -----------------------------------------------------------------------------
func RunDataflowAnalysis(config models.Config) ([]models.DataFlow, error) {
	// Configuration du logger
	logger.Setup(
		func(format string, v ...interface{}) {
			if config.Verbose {
				log.Printf("[INFO] "+format+"\n", v...)
			}
		},
		func(format string, v ...interface{}) {
			if config.Verbose {
				log.Printf("[WARNING] "+format+"\n", v...)
			}
		},
		func(format string, v ...interface{}) {
			log.Printf("[ERROR] "+format+"\n", v...)
		},
		func(format string, v ...interface{}) {
			if config.Debug {
				log.Printf("[DEBUG] "+format+"\n", v...)
			}
		},
	)

	dataflowInitial := dataFlowService.CreateDataflowInitial(config)

	// Vérifier que la variable est définie
	if config.Variable == "" {
		logger.PrintError("Variable is not defined in config")
		return dataflowInitial, nil
	}

	// Définir la langue globale
	models.GlobalLanguage = config.Language
	logger.PrintDebug("Global language set to: %s", config.Language)

	// Lire le contenu du fichier
	content, err := os.ReadFile(config.FilePath)
	if err != nil {
		logger.PrintError("error reading file: %v", err)
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Parser le contenu en syntax tree
	tree := languageService.ParseContent(content, config.Language)
	if tree == nil {
		logger.PrintError("failed to parse the file into a syntax tree")
		return nil, fmt.Errorf("failed to parse the file into a syntax tree")
	}

	logger.PrintDebug("variable selected: %s", config.Variable)

	// Variables à suivre
	variablesToTrack := map[string]bool{config.Variable: true}

	// Analyse de la variable dans la fonction
	root := tree.RootNode()
	startingFunction := nodeService.FindFunctionByLine(root, uint32(config.StartLine), config.Language)
	if startingFunction == nil {
		return dataflowInitial, nil
	}

	visitedLines := make(map[uint32]bool)
	visitedFunctions := make(map[string]*models.VisitInfo)

	// Lancer l'analyse du flux de données
	result := crawler.CrawlFromLine(root, startingFunction, content, variablesToTrack, uint32(config.StartLine), true, visitedLines, visitedFunctions)

	// Supprimer les étapes en double
	result = dataFlowService.RemoveDuplicateDataFlowStep(result, uint32(config.StartLine), config.Variable)

	// Ajouter une vérification pour les variables globales
	result = nodeService.AddGlobalVariableSteps(result, root, content, uint32(config.StartLine))

	// Afficher le résultat si le mode verbose est activé
	if config.Verbose {
		models.PrintDataFlow(result)
	}

	// Créer le modèle de dataflow final
	dataflow := dataFlowService.CreateDataflow(result, content, config.StartLine, config.Language, config.FilePath, config.Variable)

	return dataflow, nil
}
