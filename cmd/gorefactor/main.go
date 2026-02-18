package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ingvagabund/gorefactor/pkg/duplicates"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <go-file>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Find duplicated string literals
	result, err := duplicates.FindInFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	// Print results
	printDuplicates(result.Duplicates)
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
