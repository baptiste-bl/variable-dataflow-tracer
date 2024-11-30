// Fonction pour transformer le texte
function TransformText(text) {
    let text = text.toUpperCase();  // Conversion en majuscules
    let prefix = "Prefix: ";
    let finalText = AddPrefix(modifiedText, prefix);  // Appel de la fonction AddPrefix
    return finalText;
}

// Fonction pour ajouter un préfixe au texte
function AddPrefix(text, prefix) {
    return prefix + text;
}

// Fonction de test
function test() {
    let inputText = "Hello, World!";
    let result = TransformText(inputText);
    console.log(result);
}

// Fonction principale
function main() {
    test();
}

// Appel de la fonction principale
main();