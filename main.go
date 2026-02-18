package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
)

// StringLiteral represents a string literal and its location
type StringLiteral struct {
	Value string
	Line  int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <go-file>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Create a new token file set
	fset := token.NewFileSet()

	// Parse the Go file
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	// Collect all string literals
	literals := collectStringLiterals(fset, node)

	// Find duplicates
	duplicates := findDuplicates(literals)

	// Print results
	printDuplicates(duplicates)
}

// collectStringLiterals walks the AST and collects all string literals
func collectStringLiterals(fset *token.FileSet, node *ast.File) []StringLiteral {
	var literals []StringLiteral

	ast.Inspect(node, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok {
			if lit.Kind == token.STRING {
				position := fset.Position(lit.Pos())
				literals = append(literals, StringLiteral{
					Value: lit.Value,
					Line:  position.Line,
				})
			}
		}
		return true
	})

	return literals
}

// findDuplicates identifies string literals that appear more than once
func findDuplicates(literals []StringLiteral) map[string][]int {
	// Map from string value to list of line numbers
	literalMap := make(map[string][]int)

	for _, lit := range literals {
		literalMap[lit.Value] = append(literalMap[lit.Value], lit.Line)
	}

	// Filter to only keep duplicates
	duplicates := make(map[string][]int)
	for value, lines := range literalMap {
		if len(lines) > 1 {
			sort.Ints(lines)
			duplicates[value] = lines
		}
	}

	return duplicates
}

// printDuplicates prints the duplicated string literals and their locations
func printDuplicates(duplicates map[string][]int) {
	if len(duplicates) == 0 {
		fmt.Println("No duplicate string literals found.")
		return
	}

	fmt.Println("Duplicate string literals found:")
	fmt.Println()

	// Sort keys for consistent output
	var keys []string
	for key := range duplicates {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, value := range keys {
		lines := duplicates[value]
		fmt.Printf("String: %s\n", value)
		fmt.Printf("  Found %d times at lines: %v\n", len(lines), lines)
		fmt.Println()
	}
}
