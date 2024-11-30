package main

import (
	"fmt"
	"strings"
)

// Correction passage de variable avec changement de nom

func TransformText(text string) string {
	text := strings.ToUpper(text)
	prefix := "Prefix: "
	finalText := AddPrefix(text, prefix)
	return finalText
}

func AddPrefix(text, prefix string) string {
	return prefix + text
}

func test() {
	inputText := "Hello, World!"
	result := TransformText(inputText)
	fmt.Println(result)
}

func main() {
	test()
}
