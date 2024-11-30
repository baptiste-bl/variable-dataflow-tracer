import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Main {

    // Fonction DataFlowTest
    public static String DataFlowTest(String filePath, String test) {
        filePath = "example backward";
        String newPath = filePath;
        String result;
        newPath = functionTest();

        // Vérifie si le fichier existe
        File file = new File(newPath);
        if (!file.exists()) {
            result = "File does not exist";
        } else {
            // Lis le contenu du fichier
            try {
                result = new String(Files.readAllBytes(Paths.get(newPath)));
            } catch (IOException e) {
                result = "Error reading file";
            }
        }

        newPath = "test";
        return result;
    }

    // Fonction functionTest
    public static String functionTest() {
        return "example backward";
    }

    // Fonction TEST2
    public static String TEST2(String test) {
        test = "example testAAA";
        return test;
    }

    // Fonction test
    public static void test() {
        String filePath = "example.txt";
        if (filePath.isEmpty()) {
            System.out.println("File does not exist");
        }

        String testStr = "test";
        TEST2(filePathModified);

        String filePathModified = filePathModified + "1";
        String test = "1";
        String message = DataFlowTest(filePathModified, test);

        System.out.println(message);
    }

    // Fonction principale
    public static void main(String[] args) {
        String filePath = "example backward";
        test();
    }
}
