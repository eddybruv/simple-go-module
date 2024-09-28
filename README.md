# Simple Go Module

This repository contains a simple Go module demonstrating how to structure and use packages across multiple folders. The `hello` package in the `/hello` directory calls the `Hellos` function from the `greetings` package located in the `/greetings` directory.

## Folder Structure


```
/greetings
    ├── go.mod 
    ├── greetings.go 
    ├── greetings_test.go

/hello 
    ├── go.mod 
    ├── hello.go
```

### `/greetings`

The `/greetings` directory contains the `greetings` package, which defines the `Hellos` function. This function is responsible for returning greetings.

- **go.mod**: Module definition for the `greetings` package.
- **greetings.go**: Contains the `Hellos` function.
- **greetings_test.go**: Contains unit tests for the `greetings` package.

### `/hello`

The `/hello` directory contains the `main` package, which calls the `Hellos` function from the `greetings` package.

- **go.mod**: Module definition for the `main` package.
- **hello.go**: Main package that imports and calls the `Hellos` function.

## Usage

1. Navigate to the `/hello` directory.
2. Run the `hello.go` file to see the output of the `Hellos` function.

```bash
go run hello.go
