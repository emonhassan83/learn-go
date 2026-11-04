package main

import "fmt"

// practiceSliceCapacity demonstrates array vs slice, length, capacity, and reallocation behavior in Go.
func practiceSliceCapacity() {
	fmt.Println("\n--- Slice Capacity & Growth Practice ---")

	// =========================================================================
	// 📌 Example 1: Length (len) vs Capacity (cap) using make()
	// len() = বর্তমানে কয়টি উপাদান আছে
	// cap() = মেমোরির আন্ডারলাইং অ্যারেতে সর্বোচ্চ কয়টি উপাদান ধারণ করার জায়গা আছে
	// =========================================================================
	fmt.Println("\nExample 1: make([]T, len, cap)")
	// ৩টি এলিমেন্ট (initial 0), কিন্তু মেমোরিতে ৫টি এলিমেন্টের ক্যাপাসিটি বরাদ্দ
	s := make([]int, 3, 5)
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))

	// =========================================================================
	// 📌 Example 2: How capacity grows dynamically when using append()
	// যখন স্লাইসের len() তার cap() অতিক্রম করে, তখন Go স্বয়ংক্রিয়ভাবে
	// একটি নতুন ও বড় আন্ডারলাইং অ্যারে তৈরি করে (সাধারণত capacity দ্বিগুণ হয়)।
	// =========================================================================
	fmt.Println("\nExample 2: Dynamic Capacity Growth with append()")
	var numbers []int
	fmt.Println("Tracking growth as we append elements:")
	for i := 1; i <= 9; i++ {
		numbers = append(numbers, i)
		fmt.Printf("Added: %d -> Length: %d, Capacity: %d\n", i, len(numbers), cap(numbers))
	}

	// =========================================================================
	// 📌 Example 3: Sub-slicing an Array & Capacity Calculation
	// স্লাইসের capacity নির্ভর করে কোথা থেকে স্লাইস শুরু হয়েছে এবং
	// মূল অ্যারের শেষ প্রান্ত কত দূরে।
	// সূত্র: cap = (মূল অ্যারের দৈর্ঘ্য) - (স্লাইস শুরুর ইনডেক্স)
	// =========================================================================
	fmt.Println("\nExample 3: Sub-slicing an Array and Capacity Calculation")
	originalArray := [6]int{10, 20, 30, 40, 50, 60}
	fmt.Printf("Original Array: %v (Length: %d)\n", originalArray, len(originalArray))

	// ইনডেক্স 1 থেকে 4 (4 বাদ) পর্যন্ত স্লাইস
	subSlice := originalArray[1:4] // উপাদান: [20, 30, 40]
	fmt.Printf("SubSlice (arr[1:4]): %v\n", subSlice)
	fmt.Printf("SubSlice Length  : %d (4 - 1 = 3)\n", len(subSlice))
	fmt.Printf("SubSlice Capacity: %d (6 - 1 = 5)\n", cap(subSlice))

	// =========================================================================
	// 📌 Example 4: Modifying SubSlice affects Original Array (Shared Memory)
	// সাব-স্লাইস একই আন্ডারলাইং অ্যারে পয়েন্ট করে থাকে
	// =========================================================================
	fmt.Println("\nExample 4: Shared Memory between Slice and Array")
	subSlice[0] = 999 // subSlice[0] হলো originalArray[1]
	fmt.Printf("After modifying subSlice[0] = 999:\n")
	fmt.Printf("SubSlice: %v\n", subSlice)
	fmt.Printf("Original Array: %v\n", originalArray)

	// =========================================================================
	// 📌 Example 5: Creating an Independent Slice using copy()
	// মূল অ্যারের রেফারেন্স বা অতিরিক্ত মেমোরি ছেড়ে দিতে চাইলে copy() ব্যবহার করা হয়
	// =========================================================================
	fmt.Println("\nExample 5: Independent Slice using copy()")
	independentSlice := make([]int, len(subSlice))
	copy(independentSlice, subSlice)

	independentSlice[0] = 555
	fmt.Printf("Independent Slice modified: %v\n", independentSlice)
	fmt.Printf("Original Array remains unchanged: %v\n", originalArray)
}

