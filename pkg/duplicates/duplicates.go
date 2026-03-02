package duplicates

import (
	"fmt"
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

// DuplicateGroup represents a group of duplicated string literals with a suggestion
type DuplicateGroup struct {
	Value      string `json:"value"`
	Lines      []int  `json:"lines"`
	Suggestion string `json:"suggestion"`
}

// Result contains the duplicated string literals found in a file
type Result struct {
	FilePath   string          `json:"file_path"`
	Duplicates []DuplicateGroup `json:"duplicates"`
}

// FindInFile parses a Go file and finds all duplicated string literals.
// minCount sets the minimum number of occurrences to report (clamped to 2 if lower).
func FindInFile(filePath string, minCount int) (*Result, error) {
	if minCount < 2 {
		minCount = 2
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	literals := collectStringLiterals(fset, node)
	duplicates := findDuplicates(literals, minCount)

	return &Result{
		FilePath:   filePath,
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

// findDuplicates groups literals by value, keeps those meeting minCount,
// and returns them sorted by value for deterministic output.
func findDuplicates(literals []StringLiteral, minCount int) []DuplicateGroup {
	literalMap := make(map[string][]int)
	for _, lit := range literals {
		literalMap[lit.Value] = append(literalMap[lit.Value], lit.Line)
	}

	groups := []DuplicateGroup{}
	for value, lines := range literalMap {
		if len(lines) >= minCount {
			sort.Ints(lines)
			groups = append(groups, DuplicateGroup{
				Value:      value,
				Lines:      lines,
				Suggestion: generateSuggestion(value, len(lines)),
			})
		}
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Value < groups[j].Value
	})
	return groups
}

// generateSuggestion creates a human-readable suggestion message for a duplicate group
func generateSuggestion(value string, count int) string {
	return fmt.Sprintf(`String %s appears %d times. Consider extracting to a package-level constant.`, value, count)
}
