<?php

// Fonction principale
function main() {
    // numbers n'est pas défini, donc simuler l'initialisation à 1
    $numbers = 1;
    CrawlFromLine($numbers);
}

// CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées)
function CrawlFromLine($line) {
    echo 'Analyzing line: $line\n';

    // Condition pour simuler une fin de fichier à la ligne 10
    if ($line > 10) {
        echo 'End of file reached.\n';
        return;
    }

    // Simuler un appel récursif à une fonction interne
    AnalyzeFunction($line + 1);
}

// AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
function AnalyzeFunction($line) {
    echo 'Entering function at line: $line\n';

    // Simuler un appel récursif à CrawlFromLine
    CrawlFromLine($line + 1);
}

// Appel de la fonction principale
main();
?>
