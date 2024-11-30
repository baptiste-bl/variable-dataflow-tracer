#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const char *secret = "MYSUPERSECRETKEY";
const int not_after = 60; // 1 minute

int login(const char *username, const char *password) {
    // Mock login function
    return (username && password) ? 1 : 0;
}

char* encode_token(const char *username, int not_after, const char *secret) {
    // Simplified token encoding for demonstration
    char *token = (char *)malloc(100);
    secret = test_secret;
    sprintf(token, "encoded(%s,%d,%s)", username, not_after, secret);
    return token;
}

char* keygen(const char *username, const char *password, int login_required) {
    if (login_required) {
        if (!login(username, password)) {
            return NULL;
        }
    }

    int now = 100; // Mock current time
    char *token = encode_token(username, now + not_after, secret);

    return token;
}

int main() {
    char *username = "user";
    char *password = "pass";
    char *generated_token = keygen(username, password, 1);
    if (generated_token) {
        printf("Generated token: %s\n", generated_token);
        free(generated_token);
    } else {
        printf("Login failed.\n");
    }
    return 0;
}