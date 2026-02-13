# Table of Contents

- [Table of Contents](#table-of-contents)
- [Constants](#constants)
- [Basic Syntax](#basic-syntax)
- [Untyped Constants](#untyped-constants)
    - [Why is this powerful?](#why-is-this-powerful)
- [Typed Constants](#typed-constants)
- [Constant Block Declaration](#constant-block-declaration)
- [iota (Special Constant Generator)](#iota-special-constant-generator)
    - [Skipping Values](#skipping-values)
    - [Custom Start](#custom-start)
    - [Bit Shifting Example (Common in Systems Programming)](#bit-shifting-example-common-in-systems-programming)
- [What Can Be a Constant in Go?](#what-can-be-a-constant-in-go)
- [Compile-Time Nature](#compile-time-nature)
- [Difference Between const and var](#difference-between-const-and-var)
- [Important Technical Insight](#important-technical-insight)





# Constants

In Go, **constants** are values that are **fixed at compile time** and cannot be changed during program execution.

They are declared using the `const` keyword.

>[!NOTE]
> They can't use the := short declaration syntax.


---

# Basic Syntax

```go
const name type = value
```

Example:

```go
const age int = 20
const pi float64 = 3.14159
const username string = "Bilal"
```

You can also let Go infer the type:

```go
const age = 20
const pi = 3.14159
```


**Computed Constants**

Constants must be known at compile time. They are usually declared with a static value:

```go
const myInt = 15
```

However, constants can be computed as long as the computation can happen at compile time.

For example, this is valid:


```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

```go
// the current time can only be known when the program is running
const currentTime = time.Now()
```

---


# Untyped Constants

Go has something special called **untyped constants**.

Example:

```go
const x = 10
```

Here `x` is **untyped**.

This means:

* It does not get a fixed type immediately.
* It gets a type only when used.

Example:

```go
var a int = x
var b int64 = x
var c float64 = x
```

This works because `x` is untyped.

---

### Why is this powerful?

Because untyped constants are more flexible than variables.

Example:

```go
const y = 3.14
var f float32 = y
```

If `y` were a variable, this would cause a type mismatch.

---

# Typed Constants

If you explicitly declare a type:

```go
const x int = 10
```

Now `x` is strictly `int`.

You cannot assign it to another type without conversion:

```go
var a int64 = int64(x) // required conversion
```

---

# Constant Block Declaration

You can group constants:

```go
const (
    a = 1
    b = 2
    c = 3
)
```

---

# iota (Special Constant Generator)

`iota` is used inside constant blocks.

It automatically increments.

Example:

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
)
```

Output:

```
A = 0
B = 1
C = 2
```

---

### Skipping Values

```go
const (
    A = iota  // 0
    _
    C         // 2
)
```

---

### Custom Start

```go
const (
    A = iota + 1  // 1
    B             // 2
    C             // 3
)
```

---

### Bit Shifting Example (Common in Systems Programming)

```go
const (
    KB = 1 << (10 * iota)  // 1 << 0
    MB                     // 1 << 10
    GB                     // 1 << 20
    TB                     // 1 << 30
)
```

This creates:

* KB = 1
* MB = 1024
* GB = 1,048,576
* TB = 1,073,741,824

Very useful for OS-level and memory-related code (since you're into low-level stuff).

---

# What Can Be a Constant in Go?

Allowed:

* Integers
* Floats
* Complex numbers
* Strings
* Booleans

Not allowed:

* Slices
* Maps
* Arrays
* Structs
* Functions

Example (INVALID):

```go
const arr = []int{1,2,3} // ❌ Not allowed
```

---

# Compile-Time Nature

Constants:

* Must be known at compile time
* Cannot be calculated from runtime values

Invalid example:

```go
var x = 10
const y = x   // ❌ Not allowed
```

Because `x` is a variable.

---

# Difference Between const and var

| Feature                   | const | var            |
| ------------------------- | ----- | -------------- |
| Changeable                | ❌ No  | ✅ Yes          |
| Known at compile time     | ✅ Yes | ❌ Not required |
| Can be runtime calculated | ❌ No  | ✅ Yes          |
| Can be untyped            | ✅ Yes | ❌ No           |

---


# Important Technical Insight

Go constants are handled by the **compiler**, not memory.

They:

* Do not occupy runtime memory like variables
* Are substituted directly during compilation

This is why they are extremely efficient.

---
