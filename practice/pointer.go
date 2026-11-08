package main

import (
	"fmt"
)

// ============================================================================
// 📌 STRUCT DEFINITION
// ============================================================================
type Profile struct {
	ID    int
	Name  string
	Email *string // Optional/Nullable field (DB বা JSON-এ null হ্যান্ডেল করতে)
}

// UpdateNameByValue: Value receiver (কপি রিসিভ করে, তাই মূল Struct পরিবর্তন হয় না)
func (p Profile) UpdateNameByValue(newName string) {
	p.Name = newName
}

// UpdateNameByPointer: Pointer receiver (মেমোরি অ্যাড্রেস ধরে সরাসরি মিউটেট করে)
func (p *Profile) UpdateNameByPointer(newName string) {
	// 💡 Senior Guard: Go-তে nil receiver-এ মেথড কল করা সম্ভব! তাই nil চেক রাখা গুড প্র্যাকটিস
	if p == nil {
		return
	}
	p.Name = newName
}

// ============================================================================
// 🚀 practicePointers: Senior Go Engineer's Complete Guide
// ============================================================================
func practicePointers() {
	fmt.Println("==================================================================")
	fmt.Println("       🐹 GO POINTERS: CONCEPTS & SENIOR BEST PRACTICES          ")
	fmt.Println("==================================================================")

	// -------------------------------------------------------------------------
	// 📌 1. Basic Concept: Address-of (&) and Dereferencing (*) Operator
	// -------------------------------------------------------------------------
	// - & (ampersand): ভ্যারিয়েবলের RAM মেমোরি অ্যাড্রেস বের করে।
	// - * (asterisk): পয়েন্টার টাইপ ডিক্লেয়ার করতে এবং ওই অ্যাড্রেসের মূল ভ্যালু পড়তে/লিখতে (Dereference) ব্যবহৃত হয়।
	fmt.Println("\n[1] Basic Address (&) and Dereferencing (*):")
	score := 42
	var ptr *int = &score // ptr ভ্যারিয়েবলটি score-এর মেমোরি অ্যাড্রেস ধারণ করছে

	fmt.Printf("score এর ভ্যালু                : %d\n", score)
	fmt.Printf("score এর মেমোরি অ্যাড্রেস (&score): %p\n", &score)
	fmt.Printf("ptr এ জমা থাকা অ্যাড্রেস          : %p\n", ptr)
	fmt.Printf("ptr ডি-রেফারেন্স (*ptr)        : %d\n", *ptr)

	// পয়েন্টারের মাধ্যমে সরাসরি মেমোরির ভ্যালু পরিবর্তন
	*ptr = 99
	fmt.Printf("`*ptr = 99` করার পর score এর মান: %d\n", score)

	// -------------------------------------------------------------------------
	// 📌 2. Pass By Value vs Pointer in Functions
	// -------------------------------------------------------------------------
	// ⚠️ মনে রাখবেন: Go-তে "সবকিছুই Pass-by-Value"!
	// এমনকি আপনি পয়েন্টার পাঠালেও Go পয়েন্টারের (মেমোরি অ্যাড্রেসের 8-বাইটের) একটি কপি পাঠায়।
	fmt.Println("\n[2] Pass by Value vs Pointer in Functions:")
	num := 10
	fmt.Printf("শুরুতে num: %d\n", num)

	modifyByValue(num)
	fmt.Printf("modifyByValue(num) এর পর: %d (কোনো পরিবর্তন হয়নি!)\n", num)

	modifyByPointer(&num)
	fmt.Printf("modifyByPointer(&num) এর পর: %d (পরিবর্তিত হয়েছে!)\n", num)

	// -------------------------------------------------------------------------
	// 📌 3. Pointer Receiver vs Value Receiver (Struct Methods)
	// -------------------------------------------------------------------------
	// কখন Pointer Receiver ব্যবহার করবেন?
	// ১. মেথড দিয়ে Struct-এর স্টেট/ডাটা মিউটেট (পরিবর্তন) করতে চাইলে।
	// ২. Struct-এর সাইজ অনেক বড় হলে (কপি করার ওভারহেড বাঁচাতে)।
	// ৩. Consistency: কোনো Struct-এর একটি মেথড Pointer Receiver হলে বাকিগুলোও Pointer রাখা শ্রেয়।
	fmt.Println("\n[3] Pointer Receiver vs Value Receiver:")
	user := Profile{ID: 101, Name: "Alice"}

	user.UpdateNameByValue("Bob")
	fmt.Printf("UpdateNameByValue এর পর Name  : %s (চেঞ্জ হয়নি)\n", user.Name)

	user.UpdateNameByPointer("Bob")
	fmt.Printf("UpdateNameByPointer এর পর Name: %s (সরাসরি চেঞ্জ হয়েছে)\n", user.Name)

	// -------------------------------------------------------------------------
	// 📌 4. Optional / Nullable Fields (Production Pattern)
	// -------------------------------------------------------------------------
	// Go-তে primitive টাইপের শূন্য মান (0, "", false) থাকে।
	// কিন্তু ডেটাবেসের `NULL` বা API-এর অপশনাল ফিল্ড বোঝাতে পয়েন্টার (*string, *int) সেরা উপায়।
	fmt.Println("\n[4] Nullable / Optional Values:")
	userWithoutEmail := Profile{ID: 102, Name: "Hasan", Email: nil}
	if userWithoutEmail.Email == nil {
		fmt.Printf("%s-এর কোনো ইমেইল সেট করা নেই (Email is nil)\n", userWithoutEmail.Name)
	}

	validEmail := "hasan@tech.com"
	userWithEmail := Profile{ID: 103, Name: "Farhan", Email: &validEmail}
	if userWithEmail.Email != nil {
		fmt.Printf("%s-এর ইমেইল: %s\n", userWithEmail.Name, *userWithEmail.Email)
	}

	// -------------------------------------------------------------------------
	// 📌 5. Nil Pointer Safety (Panic Prevention)
	// -------------------------------------------------------------------------
	// nil পয়েন্টারকে dereference (*nilPtr) করলে রানটাইমে PANIC ঘটে!
	fmt.Println("\n[5] Nil Pointer Safety:")
	var uninitializedPtr *int // Zero-value of pointer is nil
	fmt.Printf("uninitializedPtr এর ভ্যালু: %v\n", uninitializedPtr)

	if uninitializedPtr != nil {
		fmt.Println("Value:", *uninitializedPtr)
	} else {
		fmt.Println("🛡️ Safe guard: পয়েন্টারটি nil, তাই dereference না করে panic ঠেকানো হলো।")
	}

	// -------------------------------------------------------------------------
	// 📌 6. Built-in `new()` Function
	// -------------------------------------------------------------------------
	// new(T) মেমোরিতে T টাইপের জন্য জায়গা বরাদ্দ করে, শূন্য মান দিয়ে ইনিশিয়ালাইজ করে এবং *T অ্যাড্রেস রিটার্ন করে।
	fmt.Println("\n[6] Built-in `new()` function:")
	intPtr := new(int) // equivalent to: var temp int; intPtr := &temp
	fmt.Printf("new(int) এর মেমোরি অ্যাড্রেস: %p, এর জিরো ভ্যালু: %d\n", intPtr, *intPtr)

	// -------------------------------------------------------------------------
	// 📌 7. Pointer to Pointer (Double Pointer - **T)
	// -------------------------------------------------------------------------
	fmt.Println("\n[7] Pointer to Pointer (**T):")
	val := 500
	p1 := &val  // *int -> val এর অ্যাড্রেস
	p2 := &p1   // **int -> p1 এর অ্যাড্রেস
	fmt.Printf("val: %d | *p1: %d | **p2: %d\n", val, *p1, **p2)

	// -------------------------------------------------------------------------
	// 📌 8. SENIOR GO ENGINEER'S GOLDEN RULES (কখন পয়েন্টার ব্যবহার করবেন এবং কখন করবেন না)
	// -------------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------")
	fmt.Println("💡 SENIOR ENGINEER ARCHITECTURAL INSIGHTS:")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("1. NO POINTER ARITHMETIC:")
	fmt.Println("   C/C++ এর মতো Go তে `ptr++` বা মেমোরি অফসেট লাফ দেওয়া যায় না (Memory Safety)।")
	fmt.Println()
	fmt.Println("2. SLICE, MAP, CHANNEL এ পয়েন্টার অপ্রয়োজনীয়:")
	fmt.Println("   - Slice, Map, Channel নিজেই ভেতরে মেমোরি পয়েন্টার ধারণ করে।")
	fmt.Println("   - তাই *[]int বা *map[string]string ব্যবহার করা প্রায় সবক্ষেত্রেই Anti-Pattern!")
	fmt.Println()
	fmt.Println("3. ESCAPE ANALYSIS & GC OVERHEAD (Stack vs Heap):")
	fmt.Println("   - ফাংশনের ভেতরের ভ্যারিয়েবল Stack এ থাকে (খুবই দ্রুত, কোনো GC লাগে না)।")
	fmt.Println("   - যদি ফাংশন থেকে লোকাল ভ্যারিয়েবলের পয়েন্টার রিটার্ন করা হয় (&local),")
	fmt.Println("     তাহলে Go কম্পাইলার Escape Analysis চালিয়ে তাকে Heap এ পাঠিয়ে দেয়।")
	fmt.Println("   - অতিরিক্ত Heap Allocation মানে বেশি Garbage Collection (GC) পজ।")
	fmt.Println("   - তাই ছোট struct বা primitive টাইপের জন্য পয়েন্টার না দিয়ে সরাসরি Value কপি করা পারফরম্যান্ট!")
	fmt.Println("------------------------------------------------------------------")
}

// Helper Functions
func modifyByValue(x int) {
	x = x + 100 // শুধু লোকাল কপি পরিবর্তিত হবে
}

func modifyByPointer(x *int) {
	if x != nil {
		*x = *x + 100 // মূল অ্যাড্রেসের ভ্যালু আপডেট হবে
	}
}
