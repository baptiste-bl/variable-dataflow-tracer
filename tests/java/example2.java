public class Main {

    // Fonction pour transformer le texte
    public static String TransformText(String text) {
        String text = text.toUpperCase(); // Convertir en majuscules
        String prefix = "Prefix: ";
        return AddPrefix(modifiedText, prefix);
    }

    // Fonction pour ajouter un préfixe
    public static String AddPrefix(String text, String prefix) {
        return prefix + text;
    }

    // Fonction de test
    public static void test() {
        String inputText = "Hello, World!";
        String result = TransformText(inputText);
        System.out.println(result);
    }

    // Fonction principale
    public static void main(String[] args) {
        test();
    }
}
