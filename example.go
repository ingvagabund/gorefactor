package main

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
