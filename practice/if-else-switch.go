package main

import "fmt"

// practiceIfElseSwitch demonstrates if, else if, else, and various switch statements in Go.
func practiceIfElseSwitch() {
	fmt.Println("\n--- If, Else & Switch Practice ---")

	// 1. Basic if - else
	age := 18
	if age >= 18 {
		fmt.Println("Eligible to vote!")
	} else {
		fmt.Println("Not eligible to vote!")
	}

	// 2. if with a short statement (Go's unique feature)
	// 'num' is only scoped inside this if-else block
	if num := 9; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// 3. if - else if - else ladder
	marks := 75
	if marks >= 80 {
		fmt.Println("Grade: A+")
	} else if marks >= 70 {
		fmt.Println("Grade: A")
	} else if marks >= 60 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. Basic switch statement (Note: Go automatically breaks each case, no explicit 'break' needed)
	day := "Monday"
	switch day {
	case "Friday":
		fmt.Println("Weekend is near!")
	case "Saturday", "Sunday": // Multiple matches in a single case
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a regular workday.")
	}

	// 5. Switch without condition (Tagless switch: acts as a cleaner if-else ladder)
	temperature := 28
	switch {
	case temperature > 35:
		fmt.Println("It's very hot outside.")
	case temperature >= 20 && temperature <= 35:
		fmt.Println("The weather is pleasant.")
	default:
		fmt.Println("It's cold outside.")
	}
}

