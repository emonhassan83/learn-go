package main

import "fmt"

func main() {
	// 1. Variable practice (defined in variable.go)
	// practiceVariables()

	// 2. Function practice (defined in function.go)
	// fmt.Println(SayHello("Alice"))
	// fmt.Println(makeCoffee("latte", true))
	// fmt.Println(makeCoffee("espresso", false))

	// 3. IIFE (Immediately Invoked Function Expression) Function
	makeCof := func() string {
		return makeCoffee("latte", true)
	}()

	fmt.Println(makeCof)

	// 4. If, Else & Switch practice (defined in if-else-switch.go)
	// practiceIfElseSwitch()

	// 5. For loop, break & continue practice (defined in for-loop.go)
	// practiceForLoop()

	// 6. Slice capacity practice (defined in array.go)
	// practiceSliceCapacity()

	// 7. Struct & Constructor practice (defined in struct.go)
	practiceStruct()
}
