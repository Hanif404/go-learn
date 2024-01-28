package main

import (
	"fmt"
	"hsi/generate"
)

func main() {
	// Get a greeting message and print it.
	generate1, _ := generate.GenerateNIP("ikhwan", 2024, 1, 14)
	fmt.Println("generateNIP :", generate1)
	fmt.Println("==============================================")
	generate2, _ := generate.GenerateNextNIP("ART191-00002", 5)
	fmt.Println("generateNextNIP :", generate2)
}
