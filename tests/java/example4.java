public class Main {

    // Fonction principale
    public static void main(String[] args) {
        // numbers n'est pas défini, donc simuler l'initialisation à 1
        int numbers = 1;
        CrawlFromLine(numbers);
    }

    // CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées)
    public static void CrawlFromLine(int line) {
        System.out.println("Analyzing line: " + line);

        // Condition pour simuler une fin de fichier à la ligne 10
        if (line > 10) {
            System.out.println("End of file reached.");
            return;
        }

        // Simuler un appel récursif à une fonction interne
        AnalyzeFunction(line + 1);
    }

    // AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
    public static void AnalyzeFunction(int line) {
        System.out.println("Entering function at line: " + line);

        // Simuler un appel récursif à CrawlFromLine
        CrawlFromLine(line + 1);
    }
}
