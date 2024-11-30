// Fonction pour transformer le texte
fn transform_text(text: &str) -> String {
    let text = text.to_uppercase(); // Convertir en majuscules
    let prefix = "Prefix: ";
    add_prefix(&modified_text, prefix)
}

// Fonction pour ajouter un préfixe
fn add_prefix(text: &str, prefix: &str) -> String {
    format!("{}{}", prefix, text)
}

// Fonction de test
fn test() {
    let input_text = "Hello, World!";
    let result = transform_text(input_text);
    println!("{}", result);
}

// Fonction principale
fn main() {
    test();
}
