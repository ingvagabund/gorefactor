# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gorefactor is a Go CLI tool that detects duplicated string literals in Go source files and suggests refactoring opportunities. Uses `go/ast` for parsing, zero external dependencies (stdlib only). Module path: `github.com/ingvagabund/gorefactor`, requires Go 1.24+.

## Build & Test Commands

```bash
make build                    # Build binary → ./bin/gorefactor
make test                     # Run all tests (go test -race -v ./...)
make vet                      # Run go vet
make install                  # Install to GOPATH/bin
make clean                    # Remove built binary
go test -v -run TestFindInFile/duplicates_found ./pkg/duplicates/  # Run a single test case
```

## Architecture

The codebase follows standard Go project layout with `cmd/` for the CLI entry point and `pkg/` for library code.

**Processing pipeline in `pkg/duplicates/duplicates.go`:**
1. `FindInFile(filePath, minCount)` — public API; parses a Go file into an AST via `go/parser`
2. `collectStringLiterals(fset, node)` — walks the AST with `ast.Inspect()`, extracts all `ast.BasicLit` nodes of kind `token.STRING`
3. `findDuplicates(literals, minCount)` — groups literals by value, filters to those meeting the minimum count, generates suggestions

**Key types:**
- `StringLiteral{Value, Line}` — a single occurrence (internal)
- `DuplicateGroup{Value, Lines, Suggestion}` — a group of duplicated literals with a refactoring suggestion (JSON-serializable)
- `Result{FilePath, Duplicates}` — full result for a file containing `[]DuplicateGroup` (JSON-serializable)

**CLI (`cmd/gorefactor/main.go`):** Accepts `--format text|json` and `--min-count N` flags. File path is a positional argument. JSON output writes the `Result` struct to stdout.

**Tests (`pkg/duplicates/duplicates_test.go`):** Table-driven tests using temporary files. Covers: duplicates found, no duplicates, invalid Go code, empty file, multiple different duplicates, min-count filtering, suggestion generation, JSON round-trip.

## Commit Convention

This project uses [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). Format: `type: description` (lowercase, imperative mood, no trailing period). Types: `feat`, `fix`, `docs`, `test`, `refactor`, `ci`, `chore`, `perf`. See `CONTRIBUTING.md` for full details and examples.
