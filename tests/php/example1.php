<?php
// Fonction DataFlowTest
function DataFlowTest($filePath, $test) {
    $filePath = 'example backward';
    $newPath = $filePath;
    $newPath = functionTest();

    // Vérifie si le fichier existe
    if (!file_exists($newPath)) {
        return 'File does not exist';
    } else {
        // Lis le contenu du fichier
        $content = file_get_contents($newPath);
        if ($content === false) {
            return 'Error reading file';
        }
        return $content;
    }

    $newPath = 'test';
}

// Fonction functionTest
function functionTest() {
    return 'example backward';
}

// Fonction TEST2
function TEST2($test) {
    $test = 'example testAAA';
    return $test;
}

// Fonction test
function test() {
    $filePath = 'example.txt';
    if ($filePath == '') {
        echo 'File does not exist\n';
    }

    $testStr = 'test';
    TEST2($filePathModified);

    $filePathModified = $filePathModified . '1';
    $test = 'test';
    $message = DataFlowTest($filePathModified, $test);

    echo $message . '\n';
}

// Fonction principale
function main() {
    $filePath = 'example backward';
    test();
}

main();
?>
