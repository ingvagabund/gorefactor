package duplicates

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFindInFile(t *testing.T) {
	tests := []struct {
		name           string
		code           string
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
			wantErr: false,
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
			wantErr:        false,
			wantDuplicates: map[string][]int{},
		},
		{
			name:    "invalid Go code",
			code:    `this is not valid go code`,
			wantErr: true,
		},
		{
			name: "empty file",
			code: `package main
`,
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
			wantErr: false,
			wantDuplicates: map[string][]int{
				`"foo"`: {4, 6, 8},
				`"bar"`: {5, 7},
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

			result, err := FindInFile(testFile)

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

			if len(result.Duplicates) != len(tt.wantDuplicates) {
				t.Errorf("Expected %d duplicate sets, got %d", len(tt.wantDuplicates), len(result.Duplicates))
			}

			for value, expectedLines := range tt.wantDuplicates {
				actualLines, found := result.Duplicates[value]
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
