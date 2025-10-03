package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("test-cli version %s\n", version)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "hello" {
		fmt.Println("Hello, World! 🌍")
		fmt.Println("This is a NEW feature in v0.2.0!")
		return
	}

	fmt.Println("Hello from test-cli! 🚀")
	fmt.Println("Commands:")
	fmt.Println("  test-cli version  - Show version")
	fmt.Println("  test-cli hello    - New greeting command")
}
