// Fonction principale
function main() {
    // Initialisation de l'analyse à partir de la ligne 1
    CrawlFromLine(numbers); // numbers n'est pas défini mais conservé pour correspondre au code original
}

// CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées).
function CrawlFromLine(line) {
    console.log(`Analyzing line: ${line}`);

    // Condition pour simuler une fin de fichier à la ligne 10
    if (line > 10) {
        console.log("End of file reached.");
        return;
    }

    // Simuler un appel récursif à une fonction interne
    AnalyzeFunction(line + 1);
}

// AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
function AnalyzeFunction(line) {
    console.log(`Entering function at line: ${line}`);

    // Simuler un appel récursif à CrawlFromLine
    CrawlFromLine(line + 1);
}

// Appel de la fonction principale
main();
