package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	// Parse command-line arguments
	dirnamePtr := flag.String("name", "myapp", "name of the directory to create")
	flag.Parse()

	// Define the directory structure
	dirs := []string{
		filepath.Join(*dirnamePtr, "cmd"),
		filepath.Join(*dirnamePtr, "internal", "app"),
		filepath.Join(*dirnamePtr, "internal", "config"),
		filepath.Join(*dirnamePtr, "internal", "domain", "dto"),
		filepath.Join(*dirnamePtr, "internal", "domain", "entities"),
		filepath.Join(*dirnamePtr, "internal", "domain", "repositories"),
		filepath.Join(*dirnamePtr, "internal", "handlers"),
		filepath.Join(*dirnamePtr, "internal", "middleware"),
		filepath.Join(*dirnamePtr, "internal", "routes"),
		filepath.Join(*dirnamePtr, "internal", "services"),
	}

	// Create the directory structure
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory:", dir, err)
			return
		}
	}

	// Template content for user_entities.go
	userEntityContent := `package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string ` + "`json:\"id\" gorm:\"primaryKey;type:uuid;default:uuid_generate_v4()\"`" + `
	Username string ` + "`json:\"username\"`" + `
	Password string ` + "`json:\"password\"`" + `
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return nil
}
`

	// Create the user_entities.go file
	entityFilePath := filepath.Join(*dirnamePtr, "internal", "domain", "entities", "user_entities.go")
	entityFile, err := os.Create(entityFilePath)
	if err != nil {
		fmt.Println("Failed to create file:", entityFilePath, err)
		return
	}
	defer entityFile.Close()

	// Write the template content to the file
	tmpl, err := template.New("userEntity").Parse(userEntityContent)
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}

	err = tmpl.Execute(entityFile, nil)
	if err != nil {
		fmt.Println("Failed to write to file:", entityFilePath, err)
		return
	}

	// Print success message
	fmt.Println("Directory structure created successfully, and user_entities.go file has been generated!")
}
