package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ingvagabund/gorefactor/pkg/duplicates"
)

func main() {
	var (
		format   = flag.String("format", "text", "Output format: 'text' or 'json'")
		minCount = flag.Int("min-count", 2, "Minimum number of occurrences to report (default: 2)")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <go-file>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	filePath := flag.Arg(0)

	if *format != "text" && *format != "json" {
		fmt.Fprintf(os.Stderr, "Error: format must be 'text' or 'json', got %q\n", *format)
		os.Exit(1)
	}

	result, err := duplicates.FindInFile(filePath, *minCount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	if *format == "json" {
		printJSON(result)
	} else {
		printText(result)
	}
}

// printJSON outputs the result as JSON
func printJSON(result *duplicates.Result) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}

// printText outputs the result as human-readable text
func printText(result *duplicates.Result) {
	if len(result.Duplicates) == 0 {
		fmt.Println("No duplicate string literals found.")
		return
	}

	fmt.Println("Duplicate string literals found:")
	fmt.Println()

	for _, group := range result.Duplicates {
		fmt.Printf("String: %s\n", group.Value)
		fmt.Printf("  Found %d times at lines: %v\n", len(group.Lines), group.Lines)
		fmt.Printf("  Suggestion: %s\n", group.Suggestion)
		fmt.Println()
	}
}
