**Closures in Go**

A **closure** is a **function value** that **captures variables from its surrounding scope**.

* That means the inner function can **read and modify variables defined outside of it**.
* It “closes over” the variables, hence the name **closure**.

---

### **Basic Example**

```go
func main() {
    x := 10
    adder := func(y int) int {
        x += y
        return x
    }

    fmt.Println(adder(5))  // 15
    fmt.Println(adder(3))  // 18
}
```

* `adder` is a **closure**: it can access and modify `x` even though `x` is **outside the function**.
* Every time we call `adder`, it remembers the **current value of x**.

---

### **Key points about closures in Go**

1. Closures **capture variables by reference**, not by value.
2. Variables captured by a closure **retain their state** between calls.
3. Useful for:

   * Counters
   * Accumulators
   * Deferred cleanup functions

---

### **Example: Counter Closure**

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
}
```

* `count` is **remembered** across multiple calls of the returned function.

---

### **Connection with defer / anonymous functions**

* Remember how deferred anonymous functions can read variables at execution time?
* That’s **because anonymous functions are closures**.
* Example:

```go
x := 10
defer func() {
    fmt.Println(x)  // reads latest value of x
}()
x = 20
```

* The deferred function **captures `x`** from the outer scope → prints `20`.

---

### why they can “remember” values even after the outer function has finished.

```go 
func counter() func() int {
    count := 0          // count lives in the heap for the closure
    return func() int {  // this func closes over count
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
}
```

**How it works in memory**

* Normally, local variables are on the stack and vanish when the function ends.
* But Go detects that a closure needs the variable after the function ends, so it moves that variable to the heap.
* That’s why the closure can retain the value across calls.

>[!NOTE]
> unlike cpp and c you donot have to explicitlly free the memory it is done by go run


**Summary**

* A closure is a **function that remembers variables from its surrounding scope**.
* It can **read and modify captured variables**, even after the outer function has finished.
* Widely used in Go for **counters, accumulators, callbacks, and deferred logic**.
* All anonymous functions in Go are closures.
* Regular named functions are not closures unless you pass a pointer or explicitly capture variables.

---


