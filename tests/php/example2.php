<?php

// Fonction pour transformer le texte
function TransformText($text) {
    $text = strtoupper($text); // Convertir en majuscules
    $prefix = 'Prefix: ';
    return AddPrefix($modifiedText, $prefix);
}

// Fonction pour ajouter un préfixe
function AddPrefix($text, $prefix) {
    return $prefix . $text;
}

// Fonction de test
function test() {
    $inputText = 'Hello, World!';
    $result = TransformText($inputText);
    echo $result . '\n';
}

// Fonction principale
test();

?>
