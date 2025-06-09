# Lesson 02: Variables and Data Types

## Objective
In this lesson, you will learn what variables are, the different ways to declare them in Go, and the most common basic data types you will encounter.

## What is a Variable?
Think of a variable as a labeled box where you can store a piece of information. The "label" is the variable's name, and the "information" inside is its value. In Go, every box is also designed to hold only a certain *type* of information (like text, whole numbers, etc.).

## How to Declare Variables

Go provides a few ways to create (or "declare") variables.

### 1. The `var` Keyword (Explicit Type)
This is the most formal way. You state that you are creating a variable with `var`, give it a name, specify its type, and assign a value.

**Syntax:** `var name type = value`

**Example:**
```go
var courseName string = "Go Fundamentals"
```

### 2. The `var` Keyword (Type Inference)
Go is a smart language. If you assign a value when you declare a variable, Go can often figure out the type on its own. This is called *type inference*.

**Syntax:** `var name = value`

**Example:**
```go
// Go sees "Nigel Poulton" and knows it's a string
var teacherName = "Nigel Poulton"
```

### 3. The Short Declaration Operator `:=`
This is the most common, concise, and idiomatic way to declare and initialize a variable in Go.

**Important:** The `:=` operator can **only be used inside a function**.

**Syntax:** `name := value`

**Example:**
```go
// This is a shortcut that declares the variable `numberOfLessons`
// and infers its type (int) from the value 15.
numberOfLessons := 15
```

## Basic Data Types

Every variable has a data type. Here are the most fundamental ones:

-   **`string`**: Used for text. Values are enclosed in double quotes (`"`).
-   **`int`**: Used for whole numbers (e.g., -10, 0, 42).
-   **`float64`**: Used for decimal numbers (e.g., 3.14, -0.5).
-   **`bool`**: Used for `true` or `false` values.

## Zero Values
A key feature of Go is that variables declared without an explicit initial value are given a **"zero value"**. This prevents unexpected `null` or `undefined` errors.

-   `string`: `""` (an empty string)
-   `int`: `0`
-   `float64`: `0.0`
-   `bool`: `false`

## Constants
If you have a value that should never, ever change, you declare it as a constant using the `const` keyword. Trying to change a constant will result in a compiler error.

```go
const courseWebsite = "learn-go.dev"
```
