# Generics

Generics in Go allow you to write **type-independent code** — functions and data structures that work with multiple types while preserving type safety.

Generics were introduced in Go 1.18.


# Why generics are needed

Before generics, you had two options:

### Option 1 — duplicate code

```go
func SumInt(a, b int) int { return a + b }
func SumFloat(a, b float64) float64 { return a + b }
```

Bad → repetition.

---

### Option 2 — use `interface{}` (any)

```go
func Sum(a, b interface{}) interface{}
```

Problems:

* Type assertions
* Runtime errors
* No compile-time safety

---

### Generics solve this

Write once → works for many types → compile-time safe.

---

# Basic syntax

Generic functions introduce **type parameters** inside `[]`.

```go
func Print[T any](x T) {
    fmt.Println(x)
}
```

### Explanation

| Part  | Meaning                     |
| ----- | --------------------------- |
| `T`   | type parameter              |
| `any` | constraint (means any type) |

Call:

```go
Print
Print("hello") // type inference
```

---

# Type parameter inference

Go often infers types.

```go
Print("hello")
```

Compiler deduces `T = string`.

You usually don’t write `[string]`.

---

# Generic sum example

```go
func Sum[T int | float64](a, b T) T {
    return a + b
}
```

Here:

```
T must be int OR float64
```

Union constraint.

---

# Constraints (core concept)

Constraints restrict allowed types.

### `any`

```go
T any
```

No restriction.

---

### Union constraint

```go
T int | float64 | int64
```

---

### Approximation (`~`)

Very important.

```go
type MyInt int
```

If constraint:

```go
T int
```

MyInt NOT allowed.

But:

```go
T ~int
```

Means:

> Any type whose underlying type is int

---

# 6. Generic slice function

Example — contains:

```go
func Contains[T comparable](s []T, x T) bool {
    for _, v := range s {
        if v == x {
            return true
        }
    }
    return false
}
```

### Why `comparable`

Because `==` requires comparable types.

---

# 7. Built-in constraints

Important ones:

### comparable

Supports:

* `==`
* `!=`

Used for map keys, equality checks.

---

### any

Alias for `interface{}`.

---

# 8. Custom constraints (VERY IMPORTANT)

You can define your own constraint interface.

```go
type Number interface {
    int | float64 | int32
}
```

Then:

```go
func Sum[T Number](a, b T) T {
    return a + b
}
```

This is core generics design pattern.

---

# 9. Generic data structures

Huge use case.

### Generic stack

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(x T) {
    s.items = append(s.items, x)
}

func (s *Stack[T]) Pop() T {
    n := len(s.items)
    x := s.items[n-1]
    s.items = s.items[:n-1]
    return x
}
```

Usage:

```go
var s Stack[int]
s.Push(10)
```

---

# 10. Methods with generics

Receiver must include type parameter.

```go
func (s *Stack[T]) Push(x T)
```

Because Stack itself is generic.

---

# 11. Generic map wrapper example

```go
type Cache[K comparable, V any] struct {
    m map[K]V
}

func NewCache[K comparable, V any]() *Cache[K,V] {
    return &Cache[K,V]{m: make(map[K]V)}
}
```

---

# 12. Zero value problem

Generic types have zero values.

```go
var x T
```

Used when returning empty generic value.

Example:

```go
func First[T any](s []T) T {
    if len(s) == 0 {
        var zero T
        return zero
    }
    return s[0]
}
```

Important idiom.

---

# 13. Type sets (deep concept)

Constraint interface defines a **type set**.

Example:

```go
type Signed interface {
    ~int | ~int32 | ~int64
}
```

Means:

All types whose underlying types are these.

This is key generics theory.

---

# 14. Limitations of Go generics

Important:

### No specialization

Cannot write different implementation per type.

---

### No generic methods on non-generic types

This is illegal:

```go
func (s Stack) Push[T any](x T) // not allowed
```

---

### No type parameter variance (simple model)

Go intentionally simple.

---

# 15. When to use generics

Use generics when:

* Data structures (stack, queue, set)
* Algorithms (search, sort helpers)
* Utilities (min, max)
* Libraries

Do NOT use when:

* Only one type needed
* Adds complexity
* Interfaces suffice

---

# 16. Generics vs interface (very important)

### Generics

* Compile-time polymorphism
* Performance
* Type-safe containers

### Interface

* Runtime polymorphism
* Behavior abstraction
* Dependency inversion

They solve different problems.

---

# 17. Mental model (very important)

Think:

```
Generics = type abstraction
Interfaces = behavior abstraction
```

This distinction is key for advanced Go design.

---
