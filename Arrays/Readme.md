# Arrays

An **array in Go** is a **fixed-size, homogeneous sequence of elements**.

Syntax:

```go
var arr [5]int        // array of 5 integers, initialized to 0
arr := [3]string{"a","b","c"}  // array of 3 strings
```

# Key Features

1. **Fixed size**: The size `[n]` is part of the array’s type.

```go
var a [3]int
var b [4]int
// a and b have different types: [3]int vs [4]int
```

2. **Homogeneous**: All elements must be the same type.
3. **Value type**: Arrays are copied when assigned or passed to functions.
4. **Zero-initialized**: Default values if not explicitly initialized.

```go
var a [3]int      // [0, 0, 0]
var b [2]string   // ["", ""]
var c [2]bool     // [false, false]
```

---

# Array Length

You can get length using `len()`:

```go
arr := [5]int{1,2,3,4,5}
fmt.Println(len(arr))  // 5
```

Capacity = Length (arrays in Go have no extra capacity like slices).

---

# Accessing and Modifying Elements

```go
arr := [3]int{10, 20, 30}
fmt.Println(arr[0])   // 10
arr[1] = 50
fmt.Println(arr)      // [10 50 30]
```

Indexing is **zero-based**.
Accessing out-of-bounds → **runtime panic**.

---

# Iterating Over Arrays

**Classic `for` loop**

```go
arr := [3]int{1,2,3}
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

**`range` loop**

```go
for i, v := range arr {
    fmt.Println(i, v)
}

// If you only need value:
for _, v := range arr {
    fmt.Println(v)
}
```

---

# Arrays Are Value Types

When you assign or pass an array:

```go
a := [3]int{1,2,3}
b := a   // creates a copy
b[0] = 100
fmt.Println(a) // [1 2 3] → original not affected
```

**Contrast with slices**:

```go
s1 := []int{1,2,3}
s2 := s1
s2[0] = 100
fmt.Println(s1) // [100 2 3] → underlying array shared
```

---

# Passing Arrays to Functions

Arrays are **copied** when passed to functions:

```go
func modify(a [3]int) {
    a[0] = 100
}

func main() {
    arr := [3]int{1,2,3}
    modify(arr)
    fmt.Println(arr) // [1 2 3] → original not changed
}
```

To modify the original array, pass a **pointer**:

```go
func modify(a *[3]int) {
    a[0] = 100
}

func main() {
    arr := [3]int{1,2,3}
    modify(&arr)
    fmt.Println(arr) // [100 2 3]
}
```

---

# Multi-dimensional Arrays

Go supports multi-dimensional arrays:

```go
var matrix [2][3]int
matrix[0] = [3]int{1,2,3}
matrix[1] = [3]int{4,5,6}
fmt.Println(matrix)
```

Access:

```go
fmt.Println(matrix[1][2]) // 6
```

---

# Array Type in Go

Important: The **length is part of the type**:

```go
var a [3]int
var b [4]int

// a and b are different types → cannot assign directly
// b = a // ❌ compilation error
```

---

# Partial Initialization

```go
arr := [5]int{1, 2}  // remaining elements set to 0
fmt.Println(arr)      // [1 2 0 0 0]
```

Or let Go count the size automatically:

```go
arr := [...]int{1,2,3} // compiler counts 3 elements
```

---


# Best Practices

1. Use arrays only when **size is fixed and small**.
2. For dynamic collections, **use slices**.
3. Use `[...T]{...}` for automatic length.
4. Use pointers if you want functions to modify the original array.

---

**Summary**

* Arrays = fixed-size, value types, contiguous elements
* Slices = dynamic, reference-like, growable
* Arrays include length in type → important in function signatures
* Passing arrays = copy by default
* Multi-dimensional arrays = supported

---


