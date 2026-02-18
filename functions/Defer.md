### `defer` in Go

The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.


`defer` schedules a function call to run **after the surrounding function returns**.

Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers and the like.


It runs:

* After the return statement
* In **LIFO order** (Last In, First Out)

---

## Basic Example

```go
func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
```

Output:

```
Hello
World
```

The deferred call runs at the end of `main()`.

---

## LIFO Behavior

The location of a defer statement inside a function matters. The deferred call is registered at the point where defer is executed, and it will run when the function returns. If you have multiple defer statements in a single function, they are executed in last-in, first-out order (the last deferred call runs first).


```go
func main() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}
```

Output:

```
Third
Second
First
```

Deferred calls execute in reverse order.

---

## Common Use Case: Closing Resources

```go
file, _ := os.Open("file.txt")
defer file.Close()
```

This ensures the file is closed when the function exits — even if there’s an early return.

---

## Important Rule: Arguments Are Evaluated Immediately

```go
func main() {
	x := 10
	defer fmt.Println(x)
	x = 20
}
```

Output:

```
10
```

Because `x` is evaluated when `defer` runs, not when the function ends.

if you want it to be evaluated latter you should do it like 


```go 
func main() {
	x := 10
	defer func() {
		fmt.Println(x)
	}()
	x = 20
}
```

Now output will be:

```
20
```

Because the anonymous function reads x at execution time, not at defer time.

1. `defer fmt.Println(x)` → **captures the value immediately**.

   * `x` is evaluated at the time `defer` runs.
   * Changing `x` later has **no effect**.
   * Example output: `10`

2. `defer func(){ fmt.Println(x) }()` → **reads `x` later**.

   * `x` is read when the deferred function executes at the end of `main()`.
   * Changing `x` before function exits **affects output**.
   * Example output: `20`

**Key:** First evaluates arguments immediately, second evaluates inside the deferred function at execution time.


---

## normal defer vs deferred anonymous functions


| Feature                | Normal `defer`                                        | Deferred Anonymous Function                                        |
| ---------------------- | ----------------------------------------------------- | ------------------------------------------------------------------ |
| Syntax                 | `defer f(args)`                                       | `defer func() { ... }()`                                           |
| Argument Evaluation    | Immediately when `defer` is called                    | At execution time, when function runs                              |
| Variable Capture       | Captures current value of arguments                   | Reads variables later, can see updated values                      |
| Multiple Statements    | No, only a single call                                | Yes, can include multiple statements                               |
| Latest Variable Values | Cannot see changes made after defer                   | Can see changes made after defer                                   |
| Order of Execution     | LIFO with other deferred calls                        | LIFO with other deferred calls                                     |
| Use Case               | Simple cleanup, fixed arguments                       | Dynamic cleanup, logging, or complex actions                       |
| Example                | `x := 10; defer fmt.Println(x); x = 20` → prints `10` | `x := 10; defer func() { fmt.Println(x) }(); x = 20` → prints `20` |

---

## With Named Returns

```go
func example() (x int) {
	defer func() {
		x++
	}()
	return 5
}
```

Return value will be `6`.

Why?

* `return 5` sets `x = 5`
* `defer` runs and increments `x`
* Then function exits

defer runs after the return statement has assigned values, but before the function actually returns to the caller.
so that is why we get an 6 not 5


---

## Summary

* `defer` delays execution until function ends
* Executes in LIFO order
* Arguments evaluated immediately
* Commonly used for cleanup (files, locks, connections)


