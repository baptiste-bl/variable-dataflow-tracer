#include <iostream>
using namespace std;

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
    cout << "Analyzing line: " << line << endl;

    // Condition pour simuler une fin de fichier à la ligne 10
    if (line > 10) {
        cout << "End of file reached." << endl;
        return;
    }

    // Simuler un appel récursif à une fonction interne
    AnalyzeFunction(line + 1);
}

// AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
void AnalyzeFunction(int line) {
    cout << "Entering function at line: " << line << endl;

    // Simuler un appel récursif à CrawlFromLine
    CrawlFromLine(line + 1);
}
