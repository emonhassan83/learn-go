package main

import "fmt"

// practiceForLoop contains 3 easy problems and solutions using loop, break, and continue.
func practiceForLoop() {
	fmt.Println("\n--- For Loop, Break & Continue Practice ---")

	// =========================================================================
	// 📌 Problem 1: Print only even numbers between 1 and 10 (Using `continue`)
	// Problem Description:
	// 1 থেকে 10 পর্যন্ত সংখ্যা প্রিন্ট করতে হবে, কিন্তু বিজোড় (odd) সংখ্যা
	// আসলে তা স্কিপ (skip) করতে হবে।
	// =========================================================================
	fmt.Println("\nProblem 1: Even numbers (1 to 10) using 'continue'")
	for i := 1; i <= 10; i++ {
		// বিজোড় সংখ্যা হলে বাকি কোড এক্সিকিউট না করে পরবর্তী ইটারেশনে চলে যাবে
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i) // Output: 2 4 6 8 10
	}
	fmt.Println()

	// =========================================================================
	// 📌 Problem 2: Find the first number divisible by both 4 and 7 (Using `break`)
	// Problem Description:
	// 1 থেকে শুরু করে এমন প্রথম সংখ্যাটি খুঁজে বের করতে হবে যা 4 এবং 7
	// উভয় সংখ্যা দিয়েই নিঃশেষে বিভাজ্য। পাওয়ার সাথে সাথেই লুপ বন্ধ করতে হবে।
	// =========================================================================
	fmt.Println("\nProblem 2: Find first number divisible by 4 and 7 using 'break'")
	for num := 1; num <= 100; num++ {
		if num%4 == 0 && num%7 == 0 {
			fmt.Printf("Found the first matching number: %d\n", num)
			break // সংখ্যাটি পাওয়ার সাথে সাথে লুপ বন্ধ হয়ে যাবে
		}
	}

	// =========================================================================
	// 📌 Problem 3: Sum numbers, skip zeros, and stop on negative (break + continue)
	// Problem Description:
	// সংখ্যার একটি লিস্ট থেকে শুধুমাত্র পজিটিভ সংখ্যা যোগ করো।
	// সংখ্যা 0 হলে স্কিপ (continue) করো, আর কোনো নেগেটিভ সংখ্যা আসলে
	// লুপ থামিয়ে (break) ফাইনাল যোগফল দেখাও।
	// =========================================================================
	fmt.Println("\nProblem 3: Sum numbers with 'continue' (skip 0) & 'break' (stop at negative)")
	numbers := []int{5, 10, 0, 15, 0, 20, -3, 50, 100}
	totalSum := 0

	for _, val := range numbers {
		if val < 0 {
			fmt.Printf("Encountered negative number (%d). Stopping loop!\n", val)
			break // নেগেটিভ সংখ্যা পেলে সাথে সাথে লুপ বন্ধ
		}

		if val == 0 {
			fmt.Println("Encountered 0, skipping...")
			continue // 0 যোগ না করে পরের সংখ্যায় চলে যাবে
		}

		totalSum += val
		fmt.Printf("Added %d, Current Total: %d\n", val, totalSum)
	}

	fmt.Printf("Final Sum: %d\n", totalSum)
}

