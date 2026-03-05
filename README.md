# gorefactor

A Go tool that finds duplicated string literals in Go source files and suggests refactoring opportunities.

## Project Structure

```
.
├── cmd/
│   └── gorefactor/       # Main application entry point
│       └── main.go
├── pkg/
│   └── duplicates/       # Core logic for finding duplicated literals
│       ├── duplicates.go
│       └── duplicates_test.go
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
go build -o bin/gorefactor ./cmd/gorefactor
```

Run it on a Go file:
```bash
./bin/gorefactor [flags] <path-to-go-file>
```

### Flags

- `--format` - Output format: `text` (default) or `json`
- `--min-count` - Minimum number of occurrences to report (default: 2)

### Examples

Text output (default):
```bash
./bin/gorefactor cmd/gorefactor/main.go
./bin/gorefactor pkg/duplicates/duplicates.go
```

JSON output (for AI agents and tooling):
```bash
./bin/gorefactor --format json cmd/gorefactor/main.go
```

Filter by minimum occurrences:
```bash
./bin/gorefactor --min-count 3 cmd/gorefactor/main.go
```

## Makefile Targets

- `make build` - Build the gorefactor binary (default)
- `make clean` - Remove the bin/ directory
- `make test` - Run tests
- `make vet` - Run go vet
- `make install` - Install to GOPATH/bin

## Testing

Run the test suite:
```bash
make test
```

Or using go directly:
```bash
go test -v ./...
```

The test suite includes:
- Finding duplicates in test files
- Handling files with no duplicates
- Error handling for invalid files
- Error handling for invalid Go code
- Multiple duplicate detection
- Minimum count filtering
- JSON output round-trip testing

## How it works

The tool:
1. Parses the provided Go file into an AST using Go's standard library (`pkg/duplicates`)
2. Walks the AST to collect all string literals and their line numbers
3. Identifies literals that appear more than the minimum count (default: 2)
4. Generates suggestions for extracting duplicates to constants
5. Outputs the duplicated literals with their locations and suggestions (`cmd/gorefactor`)

## Example Output

### Text Format

```
Duplicate string literals found:

String: "Hello, World!"
  Found 3 times at lines: [7 8 15]
  Suggestion: String "Hello, World!" appears 3 times. Consider extracting to a package-level constant.

String: "example"
  Found 2 times at lines: [6 11]
  Suggestion: String "example" appears 2 times. Consider extracting to a package-level constant.
```

### JSON Format

```json
{
  "file_path": "example.go",
  "duplicates": [
    {
      "value": "\"Hello, World!\"",
      "lines": [7, 8, 15],
      "suggestion": "String \"Hello, World!\" appears 3 times. Consider extracting to a package-level constant."
    },
    {
      "value": "\"example\"",
      "lines": [6, 11],
      "suggestion": "String \"example\" appears 2 times. Consider extracting to a package-level constant."
    }
  ]
}
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on contributing to this project.
