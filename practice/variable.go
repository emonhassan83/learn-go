package main

import "fmt"

// practiceVariables demonstrates variable declarations, zero values, and formatting in Go.
func practiceVariables() {
	// Group variable declaration
	var (
		name string = "Alice"
		age  uint8  = 18
	)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	// Multiple variable declaration with type inference
	var city, country string = "New York", "USA"
	fmt.Printf("City: %s, Country: %s\n", city, country)

	// Zero value declaration
	var isStudent bool
	fmt.Printf("Is student: %t\n", isStudent) // Output: Is student: false

	var score float64
	fmt.Printf("Score: %.2f\n", score) // Output: Score: 0.00

	var numbers [3]int
	fmt.Printf("Numbers: %v\n", numbers) // Output: Numbers: [0 0 0]

	// Sprintf function (formats and returns a string without printing)
	message := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(message) // Output: Hello, Alice! You are 18 years old.
}

