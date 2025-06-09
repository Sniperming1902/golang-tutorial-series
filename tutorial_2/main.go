package main

import "fmt"

func main() {
	// A variable is a storage location with a specific type and an associated name.
	// Let's explore the different ways to declare variables in Go.

	// --- METHOD 1: The 'var' keyword with an explicit type ---
	// This is the most explicit way. You declare the type and then assign a value.
	var courseName string = "Go Fundamentals"
	fmt.Println("Course:", courseName)

	// --- METHOD 2: The 'var' keyword with type inference ---
	// Go can often figure out the type based on the value you assign.
	// Here, Go infers that 'teacherName' is a string.
	var teacherName = "W Min"
	fmt.Println("Teacher:", teacherName)

	// --- METHOD 3: The Short Declaration Operator `:=` ---
	// This is the most common and concise way to declare AND initialize a variable.
	// It can only be used inside a function.
	// It automatically infers the type.
	numberOfLessons := 2
	// We use Printf here for formatted printing.
	// %v is a placeholder for the value, %T is a placeholder for the type.
	fmt.Printf("Number of lessons: %v (type is %T)\n", numberOfLessons, numberOfLessons)

	// --- BASIC DATA TYPES ---

	// string: for text
	var welcomeMessage string = "Welcome to the course!"
	fmt.Println(welcomeMessage)

	// int: for whole numbers (integers)
	var courseLengthInHours int = 3
	fmt.Printf("Course length: %v hours\n", courseLengthInHours)

	// float64: for decimal numbers (floating-point)
	var progress float64 = 0.02
	fmt.Printf("Your progress: %v\n", progress)

	// bool: for true or false values
	var isCourseCompleted bool = false
	fmt.Printf("Is the course completed? %v\n", isCourseCompleted)

	// --- ZERO VALUES ---
	// If you declare a variable without giving it a value, Go assigns a "zero value".
	// This prevents "undefined" errors common in other languages.
	var defaultName string  // Zero value for string is "" (empty string)
	var defaultAge int      // Zero value for int is 0
	var defaultRating float64 // Zero value for float is 0.0
	var defaultPublished bool // Zero value for bool is false

	fmt.Printf("Default values: Name='%s', Age=%d, Rating=%.1f, Published=%t\n",
		defaultName, defaultAge, defaultRating, defaultPublished)

	// --- CONSTANTS ---
	// Constants are variables whose values cannot be changed after they are declared.
	// Use the `const` keyword.
	const courseWebsite = "learn-go.dev"
	// courseWebsite = "new-site.com" // This line would cause a compiler error!
	fmt.Println("Visit us at:", courseWebsite)
}