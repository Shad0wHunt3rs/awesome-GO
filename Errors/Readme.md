# Errors

Error handling in **Go (Golang)** is one of its most important design philosophies. Unlike many languages that use exceptions, Go uses **explicit error values**. This leads to clear, predictable, and maintainable code.

Below is a complete, structured explanation.

---

# The Basics: Errors Are Values

In Go, errors are just values that implement the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

Most functions return `(value, error)`:

```go
result, err := doSomething()
if err != nil {
    // handle error
}
```

### Example

```go
f, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

The key principle:

> Always check errors immediately after a function call.

---

# Creating Errors

## A. Using `errors.New`

```go
import "errors"

err := errors.New("something went wrong")
```

## B. Using `fmt.Errorf`

```go
err := fmt.Errorf("user %s not found", username)
```

---

# Wrapping Errors (Go 1.13+)

Since Go 1.13, error wrapping is standardized.

### Wrapping with `%w`

```go
return fmt.Errorf("failed to read config: %w", err)
```

This preserves the original error.

## 🔹 The Situation

A function fails:

```go
err := os.Open("config.json")
```

If the file doesn’t exist, Go gives you an error like:

```
no such file or directory
```

---

## 🔹 You Want to Add More Info

You want to say:

> "failed to read config"

So you write:

```go
return fmt.Errorf("failed to read config: %v", err)
```

This prints nicely, but ❌ **it destroys the original error type**.

Go now only sees text.

---

## 🔹 The Fix: Use `%w`

```go
return fmt.Errorf("failed to read config: %w", err)
```

`%w` means:

> “Keep the original error inside this new one.”

So now:

* You add extra information ✅
* You keep the original error ✅

---

## 🔹 Why That Matters

Later you can check:

```go
errors.Is(err, os.ErrNotExist)
```

And it will still work.

If you used `%v`, it would NOT work.



---

# Inspecting Wrapped Errors

Go provides:

* `errors.Is`
* `errors.As`
* `errors.Unwrap`

## A. `errors.Is`

Checks if an error matches a specific sentinel error.

```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File does not exist")
}
```

## B. `errors.As`

Used for type assertion on wrapped errors.

```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("Path error:", pathErr.Path)
}
```

---

# Sentinel Errors

Predefined error variables.

Example from the standard library:

```go
var ErrNotExist = errors.New("file does not exist")
```

You can define your own:

```go
var ErrInvalidUser = errors.New("invalid user")
```

Then:

```go
if errors.Is(err, ErrInvalidUser) {
    // handle
}
```

---

# Custom Error Types

When you need structured error information.

```go
type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}
```

Usage:

```go
return &ValidationError{
    Field: "email",
    Msg:   "invalid format",
}
```

---

# Panic and Recover

Go does not use exceptions for normal errors. But it has:

* `panic`
* `recover`

### Panic

```go
panic("something unrecoverable happened")
```

### Recover

```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    panic("boom")
}
```

⚠️ Use panic only for:

* Programmer errors
* Truly unrecoverable states
* Initialization failures

Never use panic for normal business logic.

heres another example 

```go 
package main

import "fmt"

func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Before panic")
    panic("oh no!") // panic occurs here
    fmt.Println("After panic") // won't run
}

func main() {
    safeFunction()
    fmt.Println("Program continues normally")
}
```


---

# Error Handling Patterns

## A. Early Return Pattern (Most Common)

```go
func process() error {
    data, err := load()
    if err != nil {
        return err
    }

    if err := validate(data); err != nil {
        return err
    }

    return save(data)
}
```

This avoids deep nesting.

---

## B. Error Propagation with Context

```go
func readConfig() error {
    file, err := os.Open("config.json")
    if err != nil {
        return fmt.Errorf("opening config file: %w", err)
    }
    defer file.Close()
    return nil
}
```

Always add context when returning errors upward.

---

# Logging vs Returning Errors

Rule of thumb:

* **Library code** → return errors.
* **Application boundary (main, HTTP handler, CLI)** → log errors.

Avoid logging and returning the same error (prevents duplicate logs).

---

# Multiple Errors

Go doesn’t have built-in multi-error (before Go 1.20).

In Go 1.20+:

```go
err := errors.Join(err1, err2)
```

Check:

```go
errors.Is(err, targetErr)
```

---


# Common Anti-Patterns

**Ignoring errors**:

```go
value, _ := someFunc() // bad
```

**Comparing errors with `==` when wrapped**:

```go
if err == ErrSomething // wrong
```

Use:

```go
errors.Is(err, ErrSomething)
```

**Overusing panic**

---



# Designing Good Error Messages

Good error messages:

* Are descriptive
* Add context
* Don’t repeat unnecessary info
* Are human-readable
* Are not capitalized
* Don’t end with punctuation

Example:

```
failed to connect to database: timeout
```

---

# Philosophy of Go Error Handling

From the Go team’s philosophy:

* Explicit is better than implicit.
* Errors are part of normal flow.
* Clarity > cleverness.
* No hidden control flow (like exceptions).

This design was strongly influenced by Go’s creators including Rob Pike.

---
