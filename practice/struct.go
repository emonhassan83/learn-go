package main

import "fmt"

// 1. Defining a Struct
// Go-তে struct হলো কাস্টম ডেটা টাইপ যা বিভিন্ন ফিল্ডের সমষ্টি ধারণ করে।
type Person struct {
	Name string
	Age  int
	City string
}

// 2. Methods on Struct
// Value Receiver: struct-এর ডেটা শুধু রিড করার জন্য
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s, and I am %d years old.", p.Name, p.City, p.Age)
}

// Pointer Receiver: struct-এর ফিল্ডের মান পরিবর্তন (mutate) করার জন্য
func (p *Person) HaveBirthday() {
	p.Age++
}

// 3. Constructor Function Pattern in Go
// Go-তে কোনো ডিফল্ট class বা constructor keyword নেই।
// কনভেনশন অনুযায়ী 'New' দিয়ে শুরু হওয়া ফাংশন তৈরি করে struct-এর পয়েন্টার (*Person) রিটার্ন করা হয়।
func NewPerson(name string, age int, city string) *Person {
	// ফিল্ড ভ্যালিডেশন বা ডিফল্ট ভ্যালু সেট করার সুযোগ থাকে এখানে
	if city == "" {
		city = "Unknown"
	}

	return &Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// 4. Embedded / Nested Struct Example
type Address struct {
	Street  string
	Country string
}

type Employee struct {
	ID       int
	Person   Person  // Embedded struct
	Address  Address // Nested struct
	Position string
}

// practiceStruct demonstrates struct declaration, initialization, methods, and constructor function.
func practiceStruct() {
	fmt.Println("\n--- Struct & Constructor Practice ---")

	// =========================================================================
	// 📌 Example 1: Basic Struct Initialization
	// =========================================================================
	fmt.Println("\nExample 1: Basic Struct Initialization")

	// ক) Named Fields দিয়ে (Best practice)
	p1 := Person{
		Name: "Alice",
		Age:  24,
		City: "Dhaka",
	}
	fmt.Printf("Person 1: %+v\n", p1)

	// খ) Zero Value Struct
	var p2 Person
	fmt.Printf("Zero value Person: %+v (Name: '%s', Age: %d)\n", p2, p2.Name, p2.Age)

	// =========================================================================
	// 📌 Example 2: Using Constructor Function (NewPerson)
	// =========================================================================
	fmt.Println("\nExample 2: Using Constructor Function")

	// Constructor দিয়ে নতুন Person তৈরি
	p3 := NewPerson("Bob", 28, "Chittagong")
	fmt.Printf("Person 3 (Created via Constructor): %+v\n", p3)

	// ডিফল্ট ভ্যালু টেস্ট (ফাঁকা city দিয়ে)
	p4 := NewPerson("Charlie", 22, "")
	fmt.Printf("Person 4 (Default City check): %+v\n", p4)

	// =========================================================================
	// 📌 Example 3: Calling Methods (Value Receiver & Pointer Receiver)
	// =========================================================================
	fmt.Println("\nExample 3: Calling Methods")

	// Greet method কল করা
	fmt.Println(p3.Greet())

	// HaveBirthday কল করে বয়স বাড়ানো (Pointer receiver কাজ করবে)
	fmt.Printf("Before birthday: Age = %d\n", p3.Age)
	p3.HaveBirthday()
	fmt.Printf("After birthday : Age = %d\n", p3.Age)

	// =========================================================================
	// 📌 Example 4: Nested Struct Usage
	// =========================================================================
	fmt.Println("\nExample 4: Nested Struct (Employee)")

	emp := Employee{
		ID: 101,
		Person: Person{
			Name: "Rahim",
			Age:  30,
			City: "Sylhet",
		},
		Address: Address{
			Street:  "Zindabazar",
			Country: "Bangladesh",
		},
		Position: "Go Developer",
	}

	fmt.Printf("Employee: ID=%d, Name=%s, Position=%s, Street=%s\n",
		emp.ID, emp.Person.Name, emp.Position, emp.Address.Street)
}
