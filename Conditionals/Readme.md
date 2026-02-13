## Conditionals in Go

Conditionals allow you to make decisions in your program.

Go mainly uses:

* `if`
* `if-else`
* `if-else if-else`
* `switch`

---

# `if` Statement

### Syntax

```go
if condition {
    // code runs if condition is true
}
```

### Example

```go
age := 18

if age >= 18 {
    fmt.Println("You are an adult")
}
```

⚠ No parentheses around the condition in Go.

---

# `if-else`

```go
age := 16

if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

---

# `if-else if-else`

```go
marks := 75

if marks >= 90 {
    fmt.Println("Grade A")
} else if marks >= 75 {
    fmt.Println("Grade B")
} else {
    fmt.Println("Grade C")
}
```

---

# `if` with Initialization Statement

Go allows a short statement before the condition.

```go
if x := 10; x > 5 {
    fmt.Println("Greater than 5")
}
```

* `x` is scoped only inside the `if` block.
* Commonly used with error handling.

Example:

```go
if err := someFunction(); err != nil {
    fmt.Println("Error:", err)
}
```

This is extremely common in Go.

---

# Nested `if`

```go
age := 20
hasID := true

if age >= 18 {
    if hasID {
        fmt.Println("Entry allowed")
    }
}
```

---

# `switch` Statement

Used when you have multiple possible values.

**Basic Switch**

```go
day := 2

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```

---

**Switch Without Expression**

Go allows condition-based switch (like multiple ifs).

```go
marks := 85

switch {
case marks >= 90:
    fmt.Println("A")
case marks >= 75:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

This is cleaner than many `else if` blocks.

---

# Multiple Cases in One Line

```go
letter := "a"

switch letter {
case "a", "e", "i", "o", "u":
    fmt.Println("Vowel")
default:
    fmt.Println("Consonant")
}
```

---

# No Automatic Fallthrough

Unlike C, Go does NOT automatically fall through to the next case.

If you want it:

```go
switch x := 1; x {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
```

---
