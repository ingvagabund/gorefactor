# gorefactor

A Go tool that finds duplicated string literals in Go source files.

## Project Structure

```
.
├── cmd/
│   └── gorefactor/       # Main application entry point
│       └── main.go
├── pkg/
│   └── duplicates/       # Core logic for finding duplicated literals
│       └── duplicates.go
├── example.go            # Example file for testing
├── Makefile              # Build automation
└── go.mod
```

## Usage

Build the tool:
```bash
make build
```

Or using go directly:
```bash
go build -o gorefactor ./cmd/gorefactor
```

Run it on a Go file:
```bash
./gorefactor <path-to-go-file>
```

Example:
```bash
./gorefactor example.go
```

## Makefile Targets

- `make build` - Build the gorefactor binary (default)
- `make clean` - Remove the binary
- `make test` - Run tests
- `make install` - Install to GOPATH/bin

## How it works

The tool:
1. Parses the provided Go file into an AST using Go's standard library (`pkg/duplicates`)
2. Walks the AST to collect all string literals and their line numbers
3. Identifies literals that appear more than once (exact string match)
4. Outputs the duplicated literals with their locations (`cmd/gorefactor`)

## Example Output

```
Duplicate string literals found:

String: "Hello, World!"
  Found 3 times at lines: [6 7 12]

String: "example"
  Found 2 times at lines: [5 9]
```
