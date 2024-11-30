#include <stdio.h>
#include <math.h>

// Fonction pour calculer l'aire
double CalculateArea(double radiusTest2) {
    double area = radiusTest2;
    double test = M_PI * area * area; // Erreur conservée : 'test' n'est pas utilisé correctement
    return test;
}

// Fonction pour doubler l'aire
double DoubleArea(double area) {
    return 2 * area;
}

// Fonction pour calculer l'aire et la doubler
double CalculateAndDouble(double radiusTest) {
    double area = CalculateArea(radiusTest);
    double test = DoubleArea(test); // Erreur conservée : 'test' n'est pas défini
    double doubleArea = DoubleArea(area);
    return doubleArea;
}

// Fonction de test
void test() {
    double radius = 5.0;
    double result = CalculateAndDouble(radius);
    radius = 10.0; // Redéclaration de radius, conservée comme dans l'exemple original
    printf("%lf\n", result);
}

// Fonction principale
int main() {
    test();
    return 0;
}
