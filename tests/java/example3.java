public class Main {

    // Fonction pour calculer l'aire
    public static double CalculateArea(double radiusTest2) {
        double area = radiusTest2;
        double test = Math.PI * area * area; // Erreur conservée
        return test;
    }

    // Fonction pour doubler l'aire
    public static double DoubleArea(double area) {
        return 2 * area;
    }

    // Fonction pour calculer l'aire et la doubler
    public static double CalculateAndDouble(double radiusTest) {
        double area = CalculateArea(radiusTest);
        double test = DoubleArea(test); // Erreur conservée
        double doubleArea = DoubleArea(area);
        return doubleArea;
    }

    // Fonction de test
    public static void test() {
        double radius = 5.0;
        double result = CalculateAndDouble(radius);
        radius = 10.0; // Redéclaration de radius, conservée
        System.out.println(result);
    }

    // Fonction principale
    public static void main(String[] args) {
        test();
    }
}
