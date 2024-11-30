const secret = "MYSUPERSECRETKEY";
const not_after = 60; // 1 minute

function login(username, password) {
    // Mock login function
    return username && password;
}

function encodeToken(username, expiry, secret) {
    // Simplified token encoding for demonstration
    return `encoded(${username},${expiry},${secret})`;
}

function keygen(username, password, loginRequired = true) {
    if (loginRequired) {
        if (!login(username, password)) {
            return "";
        }
    }

    const now = 100; // Mock current time
    secret = testSecret || secret;

    return token;
}

const username = "user";
const password = "pass";
const generatedToken = keygen(username, password, true);
if (generatedToken) {
    console.log("Generated token:", generatedToken);
} else {
    console.log("Login failed.");
}
