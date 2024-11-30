import java.util.*;

public class exampleGlobal {

    public static String secret = "MYSUPERSECRETKEY";
    public static int not_after = 60; // 1 minute

    public static void main(String[] args) {
        String username = "user";
        String password = "pass";
        String generatedToken = keygen(username, password, true);
        if (generatedToken != null && !generatedToken.isEmpty()) {
            System.out.println("Generated token: " + generatedToken);
        } else {
            System.out.println("Login failed.");
        }
    }

    public static boolean login(String username, String password) {
        // Mock login function
        return username != null && !username.isEmpty() && password != null && !password.isEmpty();
    }

    public static String encodeToken(String username, int expiry, String secret) {
        secret = secret != null ? secret : "";
        // Simplified token encoding for demonstration
        return "encoded(" + username + "," + expiry + "," + secret + ")";
    }

    public static String keygen(String username, String password, boolean loginRequired) {
        if (loginRequired) {
            if (!login(username, password)) {
                return null;
            }
        }

        int now = 100; // Mock current time
        secret = secret != null ? secret : "";
        String token = encodeToken(username, now + not_after, secret);

        return token;
    }
}
