#include <iostream>
#include <string>
#include <algorithm>

using namespace std;

// Fonction pour transformer le texte
string TransformText(string text) {
    transform(text.begin(), text.end(), text.begin(), ::toupper); // Convertir en majuscules
    string prefix = "Prefix: ";
    return AddPrefix(text, prefix);
}

// Fonction pour ajouter un préfixe
string AddPrefix(string text, string prefix) {
    return prefix + text;
}

// Fonction de test
void test() {
    string inputText = "Hello, World!";
    string result = TransformText(inputText);
    cout << result << endl;
}

// Fonction principale
int main() {
    test();
    return 0;
}
