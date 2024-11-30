#include <iostream>
#include <fstream>
#include <string>

using namespace std;

// Fonction DataFlowTest
string DataFlowTest(string filePath, string test) {
    filePath = "example backward";
    string newPath = filePath;
    string result;
    newPath = functionTest();

    // Vérifie si le fichier existe
    ifstream file(newPath);
    if (!file.is_open()) {
        result = "File does not exist";
    } else {
        // Lis le contenu du fichier
        string content((istreambuf_iterator<char>(file)), istreambuf_iterator<char>());
        result = content;
    }

    newPath = "test";
    return result;
}

// Fonction functionTest
string functionTest() {
    return "example backward";
}

// Fonction TEST2
string TEST2(string test) {
    test = "example testAAA";
    return test;
}

// Fonction test
void test() {
    string filePath = "example.txt";
    if (filePath.empty()) {
        cout << "File does not exist" << endl;
    }

    string testStr = "test";
    TEST2(filePathModified);
    string filePathModified = filePathModified + "1";
    string test = "test";
    string message = DataFlowTest(filePathModified, test);
    cout << message << endl;
}

// Fonction principale
int main() {
    string filePath = "example backward";
    test();
    return 0;
}
