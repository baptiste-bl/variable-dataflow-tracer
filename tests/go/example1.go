package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func DataFlowTest(filePath string, test string) string {
	filePath := "example backward"
	newPath := filePath
	newPath := filePath
	var result string
	newPath = functionTest()

	// Vérifie si le fichier existe
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		result = "File does not exist"
	} else {
		// Lis le contenu du fichier
		content, err := ioutil.ReadFile(newPath)
		if err != nil {
			result = "Error reading file"
		} else {
			result = string(content)
		}
	}

	newPath = "test"

	return result
}

func test() {
	// Changez ceci avec le chemin du fichier que vous souhaitez tester
	filePathModified := "example.txt"
	if filePath == "" {
		fmt.Println("File does not exist")
	}
	filePath := "test"
	TEST2(filePathModified)

	filePathModified := filePathModified + 1
	test := "example.txt"
	message := DataFlowTest(filePathModified, test)

	fmt.Println(message)
}

func functionTest(filePath string) string {
	return filePath
}

func TEST2(test string) string {
	test := "example testAAA"
	return test
}

func main() {
	filePathModified := "example backward"
	test()

}
