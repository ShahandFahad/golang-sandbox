package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args holds command line arguments
	// Example go run main go sybau
	fmt.Println("Usage: go run main.go <name>")

	name := os.Args[1]
	fmt.Printf("Hello, %s!\n", name)
}
