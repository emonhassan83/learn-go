package main

import "fmt"

// SayHello is exported because it starts with a capital letter.
func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// makeCoffee prepares coffee and returns a status message.
// Since it is in package main, it can be called from main.go directly.
func makeCoffee(typeofCoffee string, isHot bool) string {
	// fmt.Printf("Making %s coffee...", typeofCoffee)
	if isHot {
		fmt.Println("It's hot!")
	} else {
		fmt.Println("It's cold!")
	}
	return fmt.Sprintf("Your %s coffee is ready!", typeofCoffee)
}
