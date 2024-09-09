package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// List of supported frameworks
var frameworks = map[string]string{
	"fiber": "github.com/gofiber/fiber/v2",
	"gin":   "github.com/gin-gonic/gin",
	"echo":  "github.com/labstack/echo/v4",
}

func main() {
	// Define command-line flags
	dirnamePtr := flag.String("name", "myapp", "Specify the name of the directory to create for the project")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nDescription:")
		fmt.Println("This tool helps to create a Go backend project structure with an optional choice of framework (fiber, gin, echo).")
		fmt.Println("It initializes the Go module, installs the chosen framework (if specified), and allows entity creation.")
	}

	// Parse command-line arguments
	flag.Parse()

	// Ask user if they want to choose a framework
	var framework string
	if askYesNo("Do you want to choose a framework? (yes/no):") {
		framework = askFramework()
	} else {
		fmt.Println("No framework chosen; proceeding without setting up a framework.")
	}

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

	// Initialize Go module and install the framework if chosen
	initGoMod(*dirnamePtr, framework)

	// Ask user if they want to create entities
	createEntities(*dirnamePtr)

	fmt.Println("Directory structure created successfully!")
	if framework != "" {
		fmt.Println("Go module has been initialized with", framework, "framework!")
	}
}

// Ask the user if they want to choose a framework
func askYesNo(prompt string) bool {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	return strings.ToLower(response) == "yes"
}

// Ask user for framework choice
func askFramework() string {
	fmt.Println("Choose a framework (fiber, gin, echo):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	framework := scanner.Text()
	if _, ok := frameworks[framework]; !ok {
		fmt.Println("Invalid framework choice. Defaulting to 'fiber'.")
		framework = "fiber"
	}
	return framework
}

// Initialize Go module with chosen framework
func initGoMod(dirname, framework string) {
	moduleName := filepath.Base(dirname)

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = dirname
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to initialize Go module:", err)
		return
	}

	if framework != "" {
		frameworkImport := frameworks[framework]
		// Add framework dependency
		cmd = exec.Command("go", "get", frameworkImport)
		cmd.Dir = dirname
		err = cmd.Run()
		if err != nil {
			fmt.Println("Failed to add framework dependency:", err)
			return
		}
	}
}

// Ask user if they want to create entities and generate them
func createEntities(dirname string) {
	if !askYesNo("Do you want to create entities? (yes/no):") {
		return
	}

	// Ask for entity details
	fmt.Println("Enter entity name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entityName := scanner.Text()

	var attributes []string
	for {
		fmt.Println("Enter attribute in the format <name> <type> (e.g., username string), or 'done' to finish:")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "done" {
			break
		}
		attributes = append(attributes, input)
	}

	// Generate entity file
	generateEntityFile(dirname, entityName, attributes)
}

// Generate entity file based on user input
func generateEntityFile(dirname, entityName string, attributes []string) {
	caser := cases.Title(language.English)
	entityTemplate := `package entities

type {{.EntityName}} struct {
{{range .Attributes}}
	{{.}}{{end}}
}`

	// Prepare template data
	attributeLines := []string{}
	for _, attr := range attributes {
		attributeParts := strings.Split(attr, " ")
		if len(attributeParts) == 2 {
			attributeLines = append(attributeLines, fmt.Sprintf("%s %s `json:\"%s\"`", caser.String(attributeParts[0]), attributeParts[1], attributeParts[0]))
		}
	}

	data := struct {
		EntityName string
		Attributes []string
	}{
		EntityName: caser.String(entityName),
		Attributes: attributeLines,
	}

	// Write entity file
	entityFilePath := filepath.Join(dirname, "internal", "domain", "entities", strings.ToLower(entityName)+"_entities.go")
	entityFile, err := os.Create(entityFilePath)
	if err != nil {
		fmt.Println("Failed to create entity file:", entityFilePath, err)
		return
	}
	defer entityFile.Close()

	tmpl, err := template.New("entity").Parse(entityTemplate)
	if err != nil {
		fmt.Println("Failed to parse entity template:", err)
		return
	}

	err = tmpl.Execute(entityFile, data)
	if err != nil {
		fmt.Println("Failed to write entity file:", err)
		return
	}

	fmt.Println("Entity file", entityFilePath, "created successfully!")
}
