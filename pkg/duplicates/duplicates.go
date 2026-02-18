package duplicates

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
)

// StringLiteral represents a string literal and its location
type StringLiteral struct {
	Value string
	Line  int
}

// Result contains the duplicated string literals found in a file
type Result struct {
	// Map from string value to list of line numbers where it appears
	Duplicates map[string][]int
}

// FindInFile parses a Go file and finds all duplicated string literals
func FindInFile(filePath string) (*Result, error) {
	// Create a new token file set
	fset := token.NewFileSet()

	// Parse the Go file
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Collect all string literals
	literals := collectStringLiterals(fset, node)

	// Find duplicates
	duplicates := findDuplicates(literals)

	return &Result{
		Duplicates: duplicates,
	}, nil
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
