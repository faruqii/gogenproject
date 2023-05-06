package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Parse command-line arguments
	dirnamePtr := flag.String("name", "myapp", "name of the directory to create")
	flag.Parse()

	// Create directory structure
	err := os.MkdirAll(filepath.Join(*dirnamePtr, "cmd"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create cmd directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "app"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create app directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "config"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create config directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "controllers"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create controllers directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "domain"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create domain directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "dto"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create dto directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "middleware"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create middleware directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "repositories"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create repositories directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "routes"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create routes directory:", err)
		return
	}
	err = os.MkdirAll(filepath.Join(*dirnamePtr, "internal", "services"), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create services directory:", err)
		return
	}

	// Print success message
	fmt.Println("Directory structure created successfully!")
}
