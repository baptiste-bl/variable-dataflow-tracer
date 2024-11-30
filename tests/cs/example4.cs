using System;

class Program
{
    // Fonction principale
    static void Main()
    {
        // numbers n'est pas défini, donc simuler l'initialisation à 1
        int numbers = 1;
        CrawlFromLine(numbers);
    }

    // CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées)
    static void CrawlFromLine(int line)
    {
        Console.WriteLine("Analyzing line: " + line);

        // Condition pour simuler une fin de fichier à la ligne 10
        if (line > 10)
        {
            Console.WriteLine("End of file reached.");
            return;
        }

        // Simuler un appel récursif à une fonction interne
        AnalyzeFunction(line + 1);
    }

    // AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
    static void AnalyzeFunction(int line)
    {
        Console.WriteLine("Entering function at line: " + line);

        // Simuler un appel récursif à CrawlFromLine
        CrawlFromLine(line + 1);
    }
}
