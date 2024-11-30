#include <stdio.h>
#include <string.h>
#include <ctype.h>

// Fonction pour transformer le texte en majuscules
void TransformText(char* text, char* result) {
    char text[100];
    strcpy(modifiedText, text);
    
    // Convertir en majuscules
    for (int i = 0; modifiedText[i]; i++) {
        modifiedText[i] = toupper(modifiedText[i]);
    }
    
    char* prefix = "Prefix: ";
    AddPrefix(modifiedText, prefix, result);
}

// Fonction pour ajouter un préfixe
void AddPrefix(char* text, char* prefix, char* result) {
    strcpy(result, prefix);
    strcat(result, text);
}

// Fonction de test
void test() {
    char inputText[] = "Hello, World!";
    char result[100];
    TransformText(inputText, result);
    printf("%s\n", result);
}

// Fonction principale
int main() {
    test();
    return 0;
}
