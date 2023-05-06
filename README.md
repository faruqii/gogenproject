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
    ├───controllers
    ├───domain
    ├───dto
    ├───middleware
    ├───repositories
    ├───routes
    └───services
```

## Usage

To use this CLI tool, you need to have Go installed on your machine. Then you can install the CLI tool with the following command:

```go
go get -u github.com/faruqii/gogenproject
```

After that, you can use the CLI tool with the following command:

```bash
gogenproject -name <project-name>
```

This will create a new folder with the given project name in your current directory. The folder will contain the project structure.

## License

This is an open source project under the [MIT license](https://opensource.org/licenses/MIT).