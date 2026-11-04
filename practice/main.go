package main

import "fmt"

func main() {
	// 1. Variable practice (defined in variable.go)
	practiceVariables()

	// 2. Function practice (defined in function.go)
	fmt.Println(SayHello("Alice"))
	fmt.Println(makeCoffee("latte", true))
	fmt.Println(makeCoffee("espresso", false))
}
