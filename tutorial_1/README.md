# Lesson 01: Hello, World!

## Objective
In this lesson, we will write our very first Go program. You will learn about the basic structure of a Go program, the `main` package, the `main` function, and how to print text to the console.

## The Code

Here is the full code for our "Hello, World!" program.

`main.go`
```go
// Every Go program starts with a package declaration.
// The `main` package is special; it tells the Go compiler that the program
// is executable (it's a program you can run).
package main

// The `import` keyword is used to include code from other packages.
// The "fmt" package (short for format) provides functions for formatting
// and printing output, like printing text to the console.
import "fmt"

// The `main` function is also special. When you run a program in the `main`
// package, execution begins in the `main` function.
func main() {
    // We call the `Println` function from the `fmt` package
    // to print a line of text to the screen.
    fmt.Println("Hello, World!")
}