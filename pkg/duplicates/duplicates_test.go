package duplicates

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestFindInFile(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		minCount       int
		wantErr        bool
		wantDuplicates map[string][]int
	}{
		{
			name: "duplicates found",
			code: `package main

import "fmt"

func example() {
	name := "example"
	fmt.Println("Hello, World!")
	message := "Hello, World!"
	fmt.Println(message)

	if name == "example" {
		fmt.Println("This is an example")
	}

	var greeting = "Hello, World!"
	fmt.Println(greeting)

	// Another use of "example"
	testName := "example"
	fmt.Println(testName)
}
`,
			minCount: 2,
			wantErr:  false,
			wantDuplicates: map[string][]int{
				`"Hello, World!"`: {7, 8, 15},
				`"example"`:       {6, 11, 19},
			},
		},
		{
			name: "no duplicates",
			code: `package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("World")
	name := "test"
}
`,
			minCount:       2,
			wantErr:        false,
			wantDuplicates: map[string][]int{},
		},
		{
			name:     "invalid Go code",
			code:     `this is not valid go code`,
			minCount: 2,
			wantErr:  true,
		},
		{
			name: "empty file",
			code: `package main
`,
			minCount:       2,
			wantErr:        false,
			wantDuplicates: map[string][]int{},
		},
		{
			name: "multiple different duplicates",
			code: `package main

func test() {
	a := "foo"
	b := "bar"
	c := "foo"
	d := "bar"
	e := "foo"
	f := "baz"
}
`,
			minCount: 2,
			wantErr:  false,
			wantDuplicates: map[string][]int{
				`"foo"`: {4, 6, 8},
				`"bar"`: {5, 7},
			},
		},
		{
			name: "min-count filtering",
			code: `package main

func test() {
	a := "foo"
	b := "bar"
	c := "foo"
	d := "bar"
	e := "foo"
	f := "baz"
}
`,
			minCount: 3,
			wantErr:  false,
			wantDuplicates: map[string][]int{
				`"foo"`: {4, 6, 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			testFile := filepath.Join(tmpDir, "test.go")

			err := os.WriteFile(testFile, []byte(tt.code), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			result, err := FindInFile(testFile, tt.minCount)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindInFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if result == nil {
				t.Fatal("Expected non-nil result")
			}

			if result.FilePath != testFile {
				t.Errorf("Expected FilePath %q, got %q", testFile, result.FilePath)
			}

			if len(result.Duplicates) != len(tt.wantDuplicates) {
				t.Errorf("Expected %d duplicate sets, got %d", len(tt.wantDuplicates), len(result.Duplicates))
			}

			// Build a map from the result for easier comparison
			resultMap := make(map[string][]int)
			for _, group := range result.Duplicates {
				resultMap[group.Value] = group.Lines
				// Verify suggestion is not empty
				if group.Suggestion == "" {
					t.Errorf("Expected non-empty suggestion for %s", group.Value)
				}
			}

			for value, expectedLines := range tt.wantDuplicates {
				actualLines, found := resultMap[value]
				if !found {
					t.Errorf("Expected to find duplicate for %s, but it was not found", value)
					continue
				}

				if !reflect.DeepEqual(actualLines, expectedLines) {
					t.Errorf("For string %s, expected lines %v, got %v", value, expectedLines, actualLines)
				}
			}
		})
	}
}

func TestGenerateSuggestion(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		count    int
		wantContain string
	}{
		{
			name:     "basic suggestion",
			value:    `"test"`,
			count:    3,
			wantContain: "appears 3 times",
		},
		{
			name:     "single occurrence",
			value:    `"foo"`,
			count:    2,
			wantContain: "appears 2 times",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suggestion := generateSuggestion(tt.value, tt.count)
			if !strings.Contains(suggestion, tt.wantContain) {
				t.Errorf("Suggestion %q should contain %q", suggestion, tt.wantContain)
			}
		})
	}
}

func TestFindInFileJSONRoundTrip(t *testing.T) {
	code := `package main

func test() {
	a := "foo"
	b := "foo"
}
`
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.go")

	err := os.WriteFile(testFile, []byte(code), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := FindInFile(testFile, 2)
	if err != nil {
		t.Fatalf("FindInFile() error = %v", err)
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}

	// Unmarshal back
	var unmarshaled Result
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal from JSON: %v", err)
	}

	// Verify round-trip
	if unmarshaled.FilePath != result.FilePath {
		t.Errorf("FilePath mismatch: got %q, want %q", unmarshaled.FilePath, result.FilePath)
	}
	if len(unmarshaled.Duplicates) != len(result.Duplicates) {
		t.Errorf("Duplicates length mismatch: got %d, want %d", len(unmarshaled.Duplicates), len(result.Duplicates))
	}
}
