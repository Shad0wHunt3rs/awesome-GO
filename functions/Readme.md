# functions in Go

## Basic Function Syntax

In Go, functions are defined using the `func` keyword.

Basic Structure

```go
func functionName(parameterName parameterType) returnType {
	// function body
}
```

Example

```go
func greet(name string) string {
	return "Hello, " + name
}
```

Calling it:

```go
message := greet("Alice")
fmt.Println(message)
```

---

# Multiple Parameters

You can pass multiple parameters.

```go
func add(a int, b int) int {
	return a + b
}
```

If parameters share the same type, you can shorten it:

```go
func add(a, b int) int {
	return a + b
}
```

---

# Multiple Return Values

Go functions can return **multiple values** — this is extremely common.

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(result)
```

This pattern (`value, error`) is idiomatic Go.

---

# Named Return Values

You can name return variables.

```go
func rectangle(width, height float64) (area float64) {
	area = width * height
	return
}
```

Here:

* `area` is declared automatically
* `return` returns it implicitly

⚠️ Use this carefully — overuse can reduce readability.

```go 
func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}
```

Is the same as:

```go
func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}
```

In the first example, x and y are the return values. At the end of the function, we could simply write return to return the values of those two variables, rather than writing return x,y

---

# Functions Without Return Values

```go
func printMessage() {
	fmt.Println("Hello")
}
```

If there’s no return type, just omit it.

---

# Ignoring Return Values

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore

For example:

```go
func getPoint() (x int, y int) {
    return 3, 4
}

// ignore y value
x, _ := getPoint()
```

Even though getPoint() returns two values, we can capture the first one and ignore the second. In Go, the blank identifier isn't just a convention; it's a real language feature that completely discards the value.

# Variadic Functions (Unlimited Arguments)

Use `...` to accept unlimited arguments.

```go
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
```

Call:

```go
sum(1, 2, 3, 4)
```

You can also pass a slice:

```go
nums := []int{1, 2, 3}
sum(nums...)
```

---

# Anonymous Functions (Function Literals)

Functions without a name.

```go
func() {
	fmt.Println("Anonymous function")
}()
```

Or assigned to a variable:

```go
multiply := func(a, b int) int {
	return a * b
}

fmt.Println(multiply(3, 4))
```

---

# Closures

A closure captures variables from its surrounding scope.

```go
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
```

Usage:

```go
c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

The inner function "remembers" `count`.

Normally, local variables live on the stack and disappear after the function returns.

But here:

The inner function is still using count.

So Go does something special:

👉 It moves count to the heap instead of the stack.

This is called escape analysis.

Because the variable "escapes" the function, Go keeps it alive.



---

# Higher-Order Functions

Functions that:

* Take functions as arguments
* Or return functions

Example:

```go
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}
```

Usage:

```go
result := applyOperation(5, 3, func(x, y int) int {
	return x + y
})
```

---

# Methods (Functions on Structs)

In Go, methods are functions with a **receiver**.

```go
type Person struct {
	Name string
}

func (p Person) greet() string {
	return "Hello, " + p.Name
}
```

Usage:

```go
p := Person{Name: "Alice"}
fmt.Println(p.greet())
```

---

# Pointer Receivers vs Value Receivers

### Value Receiver (copy)

```go
func (p Person) changeName(newName string) {
	p.Name = newName
}
```

This does NOT modify original.

### Pointer Receiver (recommended when modifying)

```go
func (p *Person) changeName(newName string) {
	p.Name = newName
}
```

This modifies the original struct.

---

# Defer in Functions

`defer` runs a function when the surrounding function returns.

```go
func example() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
```

Output:

```
Hello
World
```

Commonly used for:

* Closing files
* Unlocking mutexes
* Cleaning up resources

---

# Recursion

A function calling itself.

```go
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
```

---

# Exported vs Unexported Functions

If a function name starts with a **capital letter**, it is exported (public).

```go
func PublicFunction() {}
```

Lowercase = private to package:

```go
func privateFunction() {}
```

---

# Init Functions

Special function that runs automatically.

```go
func init() {
	fmt.Println("Package initialized")
}
```

* No parameters
* No return value
* Runs before `main()`

---

# Main Function

Program entry point:

```go
func main() {
	fmt.Println("Program starts here")
}
```

Must be in `package main`.

---

# First-Class Functions (function as data)

In Go:

* Functions are values
* Can be stored in variables
* Passed around
* Returned

```go
var operation func(int, int) int
```

>[!NOTE]
> First-Class Function is an function that is being passsed as an data (variable)

> The Higher order function is an function that is using the First-Class Function 


---

# Summary Table

| Feature                | Go Supports?                 |
| ---------------------- | ---------------------------- |
| Multiple returns       | ✅                            |
| Named returns          | ✅                            |
| Variadic               | ✅                            |
| Closures               | ✅                            |
| Higher-order functions | ✅                            |
| Methods                | ✅                            |
| Inheritance            | ❌ (uses composition instead) |

---

