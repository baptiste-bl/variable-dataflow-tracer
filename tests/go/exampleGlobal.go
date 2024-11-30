package main

import (
	"fmt"
	"time"
)

var secret = "MYSUPERSECRETKEY"
var not_after = 60 // 1 minute

func main() {
	username := "user"
	password := "pass"
	token := keygen(username, password, true)
	if token != "" {
		fmt.Println("Generated token:", token)
	} else {
		fmt.Println("No token generated.")
	}
}

func login(username, password string) bool {
	// Mock login function
	return username != "" && password != ""
}

func encodeToken(username string, nbf, exp int64, secret string) string {
	// Simplified token encoding for demonstratio
	secret = secret + username
	return fmt.Sprintf("encoded_token(username=%s, nbf=%d, exp=%d, secret=%s, algorithm=HS256)", username, nbf, exp, secret)
}

func keygen(username, password string, loginRequired bool) string {
	if loginRequired {
		if !login(username, password) {
			return ""
		}
	}

	now := time.Now().Unix()
	token := encodeToken(username, now, now+int64(not_after), secret)

	return token
}
