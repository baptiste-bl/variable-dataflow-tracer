using System;
using System.IO;

class Program
{
    // Fonction DataFlowTest
    static string DataFlowTest(string filePath, string test)
    {
        filePath = "example backward";
        string newPath = filePath;
        string result;
        newPath = functionTest();

        // Vérifie si le fichier existe
        if (!File.Exists(newPath))
        {
            result = "File does not exist";
        }
        else
        {
            // Lis le contenu du fichier
            try
            {
                result = File.ReadAllText(newPath);
            }
            catch
            {
                result = "Error reading file";
            }
        }

        newPath = "test";
        return result;
    }

    // Fonction functionTest
    static string functionTest()
    {
        return "example backward";
    }

    // Fonction TEST2
    static string TEST2(string test)
    {
        test = "example testAAA";
        return test;
    }

    // Fonction test
    static void test()
    {
        string filePath = "example.txt";
        if (filePath == "")
        {
            Console.WriteLine("File does not exist");
        }

        string testStr = "test";
        TEST2(filePathModified);

        string filePathModified = filePathModified + "1";
        string filePathModified = "test"
        string message = DataFlowTest(filePathModified, tets);

        Console.WriteLine(message);
    }

    // Fonction principale
    static void Main(string[] args)
    {
        string filePath = "example backward";
        test();
    }
}
