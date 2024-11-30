// Fonction pour calculer l'aire d'un cercle
function CalculateArea(radiusTest2) {
    let area = radiusTest2;
    test = Math.PI * area * area; // 'test' est utilisé sans être initialisé, mais conservé à l'identique
}

// Fonction pour doubler l'aire
function DoubleArea(area) {
    return 2 * area;
}

// Fonction pour calculer l'aire et la doubler
function CalculateAndDouble(radiusTest) {
    let area = CalculateArea(radiusTest);
    test = DoubleArea(test); // 'test' n'est pas défini correctement ici mais conservé pour rester fidèle à l'original
    let doubleArea = DoubleArea(area);
    return doubleArea;
}

// Fonction de test
function test() {
    let radius = 5.0;
    let result = CalculateAndDouble(radius);
    let radius = 10.0; // Redéclaration de la variable radius, conservée pour correspondre au code original
    console.log(result);
}

// Fonction principale
function main() {
    test();
}

// Appel de la fonction principale
main();
