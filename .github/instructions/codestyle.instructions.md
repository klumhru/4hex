---
applyTo: '**/*.go'
---

# Go Code Style
## Formatting
- Use `gofmt` to format your code. This is the standard formatting tool for Go and ensures that your code adheres to the Go community's conventions.
- Use `goimports` to format your code and manage imports. This tool automatically adds, removes, and sorts imports in your Go files.
- Use `golint` to check for style mistakes in your code. This tool provides suggestions for improving the readability and maintainability of your code.
- Use `go vet` to check for potential errors in your code. This tool analyzes your code and reports any issues it finds, such as unreachable code or incorrect format strings.
- Use `golangci-lint` to run multiple linters in parallel. This tool can help you catch a wide range of issues in your code, including style mistakes, potential bugs, and performance problems.
- Use `gosec` to check for security issues in your code. This tool analyzes your code for common security vulnerabilities and provides suggestions for improving security.
- Use `errcheck` to check for unhandled errors in your code. This tool analyzes your code and reports any errors that are not checked or handled properly.
- Use `staticcheck` to check for static analysis issues in your code. This tool provides a wide range of checks, including performance issues, potential bugs, and code smells.
- Use `gocyclo` to check for cyclomatic complexity in your code. This tool analyzes your code and reports the complexity of your functions, helping you identify areas that may need refactoring.
- Use `misspell` to check for common spelling mistakes in your code. This tool analyzes your code and reports any misspelled words, helping you improve the readability of your code.
- Use `go test` to run your tests. This is the standard testing tool for Go and ensures that your tests are run correctly and consistently.
- Use `go mod tidy` to clean up your Go module dependencies. This tool removes any unused dependencies and ensures that your `go.mod` file is up to date.
- Use `go doc` to generate documentation for your code. This tool generates documentation based on the comments in your code, helping you create clear and concise documentation for your project.
- Use `go generate` to run code generation tools. This tool allows you to automate the generation of code, such as mocks or other boilerplate code.
- Use `go build` to build your code. This is the standard build tool for Go and ensures that your code is built correctly and consistently.
- Use `go install` to install your code. This tool installs your code and its dependencies, making it easy to run and use your code.
- Use `go run` to run your code. This tool allows you to run your Go code directly without building it first, making it easy to test and debug your code.
- Use `go clean` to clean up your Go module. This tool removes any unnecessary files and directories, helping you keep your project organized and tidy.
- Use `go list` to list your Go module dependencies. This tool provides information about your module and its dependencies, helping you understand the structure of your project.


## Naming Conventions
- Use `CamelCase` for package names, types, and functions. This is the standard naming convention for Go and helps improve readability.
- Use `snake_case` for variable names. This is a common convention in Go and helps improve readability.
- Use `UPPER_SNAKE_CASE` for constants. This is a common convention in Go and helps improve readability.
- Use `lowerCamelCase` for struct fields and method receivers. This is a common convention in Go and helps improve readability.
- Use `PascalCase` for exported types and functions. This is a common convention in Go and helps improve readability.
- Use `lowercase` for unexported types and functions. This is a common convention in Go and helps improve readability.
- Use `short` and `meaningful` names for variables. This helps improve readability and maintainability of your code.
- Use `descriptive` names for functions and methods. This helps improve readability and maintainability of your code.
- Use `singular` names for types and `plural` names for slices. This helps improve readability and maintainability of your code.
- Use `singular` names for maps
- Use `singular` names for channels

## Test Patterns
- Use `*_test.go` for test files. This is the standard naming convention for Go test files and helps improve readability.
- Use `Test*` for test functions. This is the standard naming convention for Go test functions and helps improve readability.
- Create Data Driven Tests for testing multiple scenarios with the same logic. This helps improve code reuse and maintainability.
- Use testify/assert for test assertions.
- Use assert := assert.New(t) for test assertions.
- Use require := require.New(t) for test assertions.