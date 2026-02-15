## **Currying**

**Currying** is a technique in programming where a function with multiple parameters is transformed into a sequence of functions, each taking **one parameter at a time**.

In simpler words:

> Instead of giving all inputs to a function at once, you give them one by one, and each time you get a new function waiting for the next input.

---

## **Why Currying?**

* Makes code **more modular**.
* Allows **reusing a function with some fixed parameters**.
* Helps in **functional programming** style.

---

## **How Currying Works in Go**

Go doesn’t have built-in currying like JavaScript or Haskell, but we can implement it using **closures** (anonymous functions that “remember” the variables from their scope).



**Example 1: Simple Addition Function**

We want a function that adds two numbers:

```go
package main

import "fmt"

func add(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

func main() {
    add5 := add(5) // fix the first number as 5
    fmt.Println(add5(3)) // 5 + 3 = 8
    fmt.Println(add5(10)) // 5 + 10 = 15

    fmt.Println(add(2)(4)) // directly call with both parameters: 2 + 4 = 6
}
```

**Explanation:**

1. `add(x int)` returns **another function** `func(y int) int`.
2. The inner function “remembers” the value of `x` due to closure.
3. We can create a new function with a fixed `x` and reuse it multiple times (`add5`).

---

**Example 2: Multiplication Currying**

```go
package main

import "fmt"

func multiply(a int) func(int) int {
    return func(b int) int {
        return a * b
    }
}

func main() {
    double := multiply(2) // fix first number as 2
    triple := multiply(3) // fix first number as 3

    fmt.Println(double(5)) // 2*5 = 10
    fmt.Println(triple(7)) // 3*7 = 21
}
```

* `double` is a new function that multiplies any number by 2.
* `triple` multiplies any number by 3.
* That’s **currying in action**: partial application of arguments.

---

### **4. Step by Step Concept**

1. **Original function:**
   `f(x, y) = x + y`
2. **Curried function:**
   `f(x) => g(y) => x + y`

* You first give `x` → you get a new function `g(y)`.
* Then give `y` → you get the final result.

---

### **Real-World Use Case**

Suppose you have a **discount calculator**:

```go
package main

import "fmt"

func discountCalculator(rate float64) func(price float64) float64 {
    return func(price float64) float64 {
        return price * (1 - rate)
    }
}

func main() {
    tenPercent := discountCalculator(0.10)
    twentyPercent := discountCalculator(0.20)

    fmt.Println(tenPercent(100))  // 90
    fmt.Println(twentyPercent(100)) // 80
}
```

* We “fix” the discount rate and create specialized functions for different customers.


Yes! You can absolutely do the same thing without currying, using a normal function with both parameters at once. Currying just gives you the flexibility to “fix” some arguments and reuse the function, but it’s not mandatory.

Here’s how your example would look without currying:

```go
package main

import "fmt"

// Normal function: both rate and price are passed together
func discount(price float64, rate float64) float64 {
    return price * (1 - rate)
}

func main() {
    fmt.Println(discount(100, 0.10)) // 90
    fmt.Println(discount(100, 0.20)) // 80
}
```

---

**Key Points:**

1. Currying = splitting multi-parameter function into a chain of single-parameter functions.
2. Go achieves it using **closures**.
3. Useful for **partial application**, reusability, and functional-style programming.

---

