using System;

class Program
{
    // Fonction pour calculer l'aire
    static double CalculateArea(double radiusTest2)
    {
        double area = radiusTest2;
        double test = Math.PI * area * area; // Erreur conservée
        return test;
    }

    // Fonction pour doubler l'aire
    static double DoubleArea(double area)
    {
        return 2 * area;
    }

    // Fonction pour calculer l'aire et la doubler
    static double CalculateAndDouble(double radiusTest)
    {
        double area = CalculateArea(radiusTest);
        double test = DoubleArea(test); // Erreur conservée
        double doubleArea = DoubleArea(area);
        return doubleArea;
    }

    // Fonction de test
    static void test()
    {
        double radius = 5.0;
        double result = CalculateAndDouble(radius);
        radius = 10.0; // Redéclaration de radius, conservée
        Console.WriteLine(result);
    }

    // Fonction principale
    static void Main(string[] args)
    {
        test();
    }
}
