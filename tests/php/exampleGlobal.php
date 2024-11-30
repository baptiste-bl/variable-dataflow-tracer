<?php

$secret = 'MYSUPERSECRETKEY';
$not_after = 60; // 1 minute

function login($username, $password) {
    // Mock login function
    return !empty($username) && !empty($password);
}

function encodeToken($username, $expiry, $secret) {
    // Simplified token encoding for demonstration
    return 'encoded($username,$expiry,$secret)';
}

function keygen($username, $password, $loginRequired = true) {
    if ($loginRequired) {
        if (!login($username, $password)) {
            return '';
        }
    }

    $now = 100; // Mock current time
    $GLOBALS['secret'] = $secret;
    $token = encodeToken($username, $now + $GLOBALS['not_after'], $GLOBALS['secret']);

    return $token;
}

$username = 'user';
$password = 'pass';
$generatedToken = keygen($username, $password, true);
if (!empty($generatedToken)) {
    echo 'Generated token: $generatedToken\n';
} else {
    echo 'Login failed.\n';
}

?>