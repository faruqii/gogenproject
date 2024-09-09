# Go-Generate-Project-Structure

This is my personal project structure for Go projects. It is based on the standard Go project structure, but with some additions based on my personal preferences.

## Backstory

When i started new Go projects, i always had to create the same folder structure over and over again. So i decided to create a template repository, which i can use to generate a new project structure. So I created this CLI tool, which can be used to generate a new project structure based on this repository.
Here My personal Go project structure.
```
├───cmd
└───internal
    ├───app
    ├───config
    ├───domain
    │   ├───dto
    │   ├───entities
    │   └───repositories
    ├───handlers
    ├───middleware
    ├───routes
    └───services
```

## What's New?
- `Framework Selection`: The CLI tool now offers the option to choose a Go framework (Fiber, Gin, or Echo). The user can decide whether to set up a framework or not, making the selection process flexible and optional.
- `Go Module Initialization`: Automatically initializes a Go module (go.mod) with the project name and installs the chosen framework if specified.
- `Entity Generation`: Users can now generate custom entity files by specifying the entity name and attributes interactively.


## Usage

To use this CLI tool, you need to have Go installed on your machine. Then you can install the CLI tool with the following command:

```go
go install github.com/faruqii/gogenproject/v2@latest
```

After that, you can use the CLI tool with the following command:

```bash
gogenproject -name <project-name>
```

This will create a new folder with the given project name in your current directory. The folder will contain the project structure.

## License

This is an open source project under the [MIT license](https://opensource.org/licenses/MIT).