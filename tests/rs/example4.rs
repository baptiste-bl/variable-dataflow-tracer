// Fonction principale
fn main() {
    // numbers n'est pas défini, donc simuler l'initialisation à 1
    let numbers = 1;
    crawl_from_line(numbers);
}

// CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées)
fn crawl_from_line(line: i32) {
    println!("Analyzing line: {}", line);

    // Condition pour simuler une fin de fichier à la ligne 10
    if line > 10 {
        println!("End of file reached.");
        return;
    }

    // Simuler un appel récursif à une fonction interne
    analyze_function(line + 1);
}

// AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
fn analyze_function(line: i32) {
    println!("Entering function at line: {}", line);

    // Simuler un appel récursif à CrawlFromLine
    crawl_from_line(line + 1);
}
