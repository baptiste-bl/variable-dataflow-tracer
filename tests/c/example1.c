#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Fonction DataFlowTest
char* DataFlowTest(char* filePath, char test) {
    filePath = "example backward";
    char* newPath = filePath;
    char* result;
    newPath = functionTest();

    // Vérifie si le fichier existe
    FILE *file = fopen(newPath, "r");
    if (file == NULL) {
        result = "File does not exist";
    } else {
        // Lis le contenu du fichier
        fseek(file, 0, SEEK_END);
        long length = ftell(file);
        fseek(file, 0, SEEK_SET);
        char *content = malloc(length);
        if (content) {
            fread(content, 1, length, file);
            result = content;
        } else {
            result = "Error reading file";
        }
        fclose(file);
    }

    newPath = "test";
    return result;
}

// Fonction functionTest
char* functionTest() {
    return "example backward";
}

// Fonction TEST2
char* TEST2(char* test) {
    test = "example testAAA";
    return test;
}

// Fonction test
void test() {
    char* filePath = "example.txt";
    if (strcmp(filePath, "") == 0) {
        printf("File does not exist\n");
    }

    char* testStr = "test";
    TEST2(filePathModified);

    // Ici on simule une addition avec une chaîne de caractères
    char filePathModified[50];
    sprintf(filePathModified, "%s1", filePathModified);
    char test = "test";
    char* message = DataFlowTest(filePathModified, test);
    printf("%s\n", message);
}

// Fonction principale
int main() {
    char* filePathModified = "example backward";
    test();
    return 0;
}
