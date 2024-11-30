<?php

// Fonction pour calculer l'aire
function CalculateArea($radiusTest2) {
    $area = $radiusTest2;
    $test = pi() * $area * $area; // Erreur conservée
    return $test;
}

// Fonction pour doubler l'aire
function DoubleArea($area) {
    return 2 * $area;
}

// Fonction pour calculer l'aire et la doubler
function CalculateAndDouble($radiusTest) {
    $area = CalculateArea($radiusTest);
    $test = DoubleArea($test); // Erreur conservée
    $doubleArea = DoubleArea($area);
    return $doubleArea;
}

// Fonction de test
function test() {
    $radius = 5.0;
    $result = CalculateAndDouble($radius);
    $radius = 10.0; // Redéclaration de radius, conservée
    echo $result . '\n';
}

// Fonction principale
test();

?>
