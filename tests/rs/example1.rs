use std::fs;
use std::path::Path;

// Fonction DataFlowTest
fn data_flow_test(file_path: &str, test: &str) -> String {
    let file_path = "example backward";
    let mut new_path = file_path.to_string();
    new_path = function_test();

    // Vérifie si le fichier existe
    if !Path::new(&new_path).exists() {
        return "File does not exist".to_string();
    } else {
        // Lis le contenu du fichier
        match fs::read_to_string(new_path) {
            Ok(content) => return content,
            Err(_) => return "Error reading file".to_string(),
        }
    }
}

// Fonction functionTest
fn function_test() -> String {
    "example backward".to_string()
}

// Fonction TEST2
fn test2(test: &str) -> String {
    test = "example testAAA";
    "example testAAA".to_string()
}

// Fonction test
fn test() {
    let file_path = "example.txt";
    if file_path.is_empty() {
        println!("File does not exist");
    }

    let test_str = "test";
    test2(file_path_modified);

    let file_path_modified = format!("{}1", file_path_modified);
    let test = "test";
    let message = data_flow_test(&file_path_modified, test);

    println!("{}", message);
}

// Fonction principale
fn main() {
    let file_path = "example backward";
    test();
}
