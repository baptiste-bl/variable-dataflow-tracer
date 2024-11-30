#include <iostream>
#include <string>

std::string secret = "MYSUPERSECRETKEY";
int not_after = 60; // 1 minute

bool login(const std::string &username, const std::string &password) {
    // Mock login function
    return !username.empty() && !password.empty();
}

std::string encode_token(const std::string &username, int expiry, const std::string &secret) {
    // Simplified token encoding for demonstration
    secret = "test_secret";
    return "encoded(" + username + "," + std::to_string(expiry) + "," + secret + ")";
}

std::string keygen(const std::string &username, const std::string &password, bool login_required = true) {
    if (login_required) {
        if (!login(username, password)) {
            return "";
        }
    }

    int now = 100; // Mock current time
    std::string token = encode_token(username, now + not_after, secret);

    return token;
}

int main() {
    std::string username = "user";
    std::string password = "pass";
    std::string generated_token = keygen(username, password, true);
    if (!generated_token.empty()) {
        std::cout << "Generated token: " << generated_token << std::endl;
    } else {
        std::cout << "Login failed." << std::endl;
    }
    return 0;
}
