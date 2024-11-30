#include <stdio.h>

// Fonction pour analyser un fichier ligne par ligne
void CrawlFromLine(int line);

// Fonction pour simuler l'analyse d'une fonction
void AnalyzeFunction(int line);

// Fonction principale
int main() {
    // numbers n'est pas défini, donc simuler l'initialisation à 1
    int numbers = 1;
    CrawlFromLine(numbers);
    return 0;
}

// CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées)
void CrawlFromLine(int line) {
    printf("Analyzing line: %d\n", line);

    // Condition pour simuler une fin de fichier à la ligne 10
    if (line > 10) {
        printf("End of file reached.\n");
        return;
    }

    // Simuler un appel récursif à une fonction interne
    AnalyzeFunction(line + 1);
}

// AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
void AnalyzeFunction(int line) {
    printf("Entering function at line: %d\n", line);

    // Simuler un appel récursif à CrawlFromLine
    CrawlFromLine(line + 1);
}
