using System;

class Program
{
    // Fonction pour transformer le texte
    static string TransformText(string text)
    {
        string text = text.ToUpper(); // Convertir en majuscules
        string prefix = "Prefix: ";
        return AddPrefix(modifiedText, prefix);
    }

    // Fonction pour ajouter un préfixe
    static string AddPrefix(string text, string prefix)
    {
        return prefix + text;
    }

    // Fonction de test
    static void test()
    {
        string inputText = "Hello, World!";
        string result = TransformText(inputText);
        Console.WriteLine(result);
    }

    // Fonction principale
    static void Main(string[] args)
    {
        test();
    }
}
