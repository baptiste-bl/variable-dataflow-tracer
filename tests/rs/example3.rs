use std::f64::consts::PI;

// Fonction pour calculer l'aire
fn calculate_area(radius_test2: f64) -> f64 {
    let area = radius_test2;
    let test = PI * area * area; // Erreur conservée
    return test;
}

// Fonction pour doubler l'aire
fn double_area(area: f64) -> f64 {
    return 2.0 * area;
}

// Fonction pour calculer l'aire et la doubler
fn calculate_and_double(radius_test: f64) -> f64 {
    let area = calculate_area(radius_test);
    let test = double_area(test); // Erreur conservée
    let double_area = double_area(area);
    return double_area;
}

// Fonction de test
fn test() {
    let mut radius = 5.0;
    let result = calculate_and_double(radius);
    radius = 10.0; // Redéclaration de radius, conservée
    println!("{}", result);
}

// Fonction principale
fn main() {
    test();
}
