# Loops

Loops in **Go** are simpler than many other languages because Go has only **one looping keyword: `for`**.
There is **no `while` or `do-while`** — everything is done using `for`.


# Basic `for` Loop (C-style)

This is the most common loop format.

```go
for initialization; condition; post {
    // body
}
```

**Example**:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**How it works**:

1. `i := 0` → runs once (initialization)
2. `i < 5` → checked before every iteration
3. Loop body runs
4. `i++` → runs after each iteration

Output:

```
0
1
2
3
4
```

---

# `while`-style Loop in Go

Go doesn’t have `while`, but you can simulate it:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

This behaves exactly like:

```c
while (i < 5)
```

Only condition is written.

---

# Infinite Loop

If you omit everything:

```go
for {
    fmt.Println("Infinite")
}
```

This runs forever.

To stop it, use `break`.

Example:

```go
for {
    fmt.Println("Hello")
    break
}
```

---

# `range` Loop

Used to iterate over:

* Arrays
* Slices
* Strings
* Maps
* Channels

---

## Loop over Array/Slice

```go
nums := []int{10, 20, 30}

for index, value := range nums {
    fmt.Println(index, value)
}
```

Output:

```
0 10
1 20
2 30
```

If you only need value:

```go
for _, value := range nums {
    fmt.Println(value)
}
```

`_` is blank identifier (ignore value).

---

## Loop over String

```go
str := "Go"

for i, ch := range str {
    fmt.Println(i, ch)
}
```

Important:

* `ch` is a **rune (Unicode code point)**
* Go strings are UTF-8

If you print as character:

```go
fmt.Printf("%c\n", ch)
```

---

## Loop over Map

```go
m := map[string]int{
    "a": 1,
    "b": 2,
}

for key, value := range m {
    fmt.Println(key, value)
}
```

Important:

* Map iteration order is **not guaranteed**.

---

## Loop over Channel

```go
for value := range ch {
    fmt.Println(value)
}
```

This continues until channel is closed.

---

# `break` Statement

Stops the loop immediately.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

Output:

```
0 1 2 3 4
```

---

# `continue` Statement

Skips current iteration.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

Output:

```
0 1 3 4
```

---

# Labeled Loops

Go supports labels.

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer
        }
        fmt.Println(i, j)
    }
}
```

This breaks the **outer loop**, not just inner loop.

You can also use:

```go
continue outer
```

---

# Nested Loops

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Println(i, j)
    }
}
```

Used for:

* Matrices
* Grids
* Game boards

---

# Loop Variable Scope

Important concept:

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

`i` exists only inside the loop.

If you define outside:

```go
i := 0
for ; i < 3; i++ {
}
fmt.Println(i)  // accessible
```

---

# Modifying Loop Variables Inside `range`

Common mistake:

```go
nums := []int{1,2,3}
for _, v := range nums {
    v = 10   // does NOT change original slice
}
```

Why?
Because `v` is a copy.

Correct way:

```go
for i := range nums {
    nums[i] = 10
}
```

---

# Performance Considerations

* `range` over slice → efficient
* `range` over string → decodes UTF-8 (slightly slower)
* Map iteration → randomized order
* Avoid unnecessary allocations inside loops

---
