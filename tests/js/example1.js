const fs = require('fs');

// Fonction pour tester le flux de données
function DataFlowTest(filePath, test) {
    filePath = "example backward";
    let newPath = filePath;
    let result = '';
    newPath = functionTest();

    // Vérifie si le fichier existe
    if (!fs.existsSync(newPath)) {
        result = "File does not exist";
    } else {
        try {
            // Lis le contenu du fichier
            const content = fs.readFileSync(newPath, 'utf8');
            result = content;
        } catch (err) {
            result = "Error reading file";
        }
    }

    newPath = "test";

    return result;
}

// Fonction d'exemple
function functionTest(filePath) {
    return filePath;
}

// Fonction de test
function test() {
    // Changez ceci avec le chemin du fichier que vous souhaitez tester
    let filePath = "example.txt";
    if (filePath === "") {
        console.log("File does not exist");
    }

    let test = "test";
    TEST2(filePathModified);

    filePathModified = filePathModified + 1;    
    // Ceci va produire une erreur car filePath est une chaîne de caractères et "filePath + 1" n'est pas valide
    test = "test"
    let message = DataFlowTest(filePathModified, test);

    console.log(message);
}

// Deuxième définition de la fonction (la redéclaration est conservée)
function functionTest(filePath) {
    return filePath;
}

// Fonction TEST2
function TEST2(test) {
    test = "example testAAA";
    return test;
}

// Fonction principale
function main() {
    let filePath = "example backward";
    test();
}

// Appel de la fonction principale
main();
