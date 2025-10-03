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
    
    fmt.Println("Hello from test-cli! 🚀")
    fmt.Println("Run 'test-cli version' to see the version")
}	
