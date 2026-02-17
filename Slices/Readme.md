## Slices in Go

In Go, **slices** are a flexible, powerful abstraction built on top of arrays. They are one of the most commonly used data structures in Go.

A **slice** is a dynamically-sized, flexible view into the elements of an underlying array.

Unlike arrays:

* Arrays have **fixed size**
* Slices can **grow and shrink**
* Slices are used much more frequently than arrays in real-world Go programs

---

# Slice vs Array

### Array

```go
var arr [5]int
```

* Size is part of the type (`[5]int`)
* Cannot change size
* Value type (copied on assignment)

### Slice

```go
var s []int
```

* No size specified
* Dynamic length
* Reference type (points to underlying array)

---

# Internal Structure of a Slice

A slice is a small struct containing:

```go
type slice struct {
    ptr *T   // pointer to underlying array
    len int  // length
    cap int  // capacity
}
```

### Important Terms:

* **Length (len)** → Number of elements currently in slice
* **Capacity (cap)** → Number of elements slice can grow to before reallocating

Example:

```go
s := make([]int, 3, 5)
```

* len = 3
* cap = 5

---

# Creating Slices

## 1. Using `make`

```go
s := make([]int, 3)        // len=3, cap=3
s := make([]int, 3, 5)     // len=3, cap=5
```

## 2. Slice Literal

```go
s := []int{1, 2, 3}
```

## 3. From an Array

```go
arr := [5]int{1,2,3,4,5}
s := arr[1:4]
```

This creates a slice:

* Includes index 1,2,3
* Excludes index 4




1. Arrays have a fixed size, which makes them less flexible than slices.
2. Slicing an array can unintentionally share memory and cause side effects.
3. Most Go code and APIs use slices (`[]T`), so creating slices directly is more idiomatic.


```go
arr := [3]int{1,2,3}
s := arr[:]
```

Memory:

```bash
arr: [1 2 3]
s  ──┘
```

They share the same underlying array. so if you do any chane in the slice it will also affect the array 

but what if we have to append or exclude so in that case 

```go 
s = append(s, 4)
```

So Go does this:

* Allocate new bigger array
* Copy [1 2 3]
* Add 4
* Make s point to new array

Now memory looks like:

```bash
arr:        [1 2 3]        (unchanged)

new array:  [1 2 3 4]
              ↑
              s
```

so now the arr and slice are not connected a change in slice will not effect the arr 

this is why In real-world Go code, developers almost always use make or slice literals instead of slicing an array directly.

---

# Slicing Syntax

```go
s[a:b]
```

Meaning:

* Start at index `a`
* End before index `b`

>[!NOTE]
> the first no is inclusive & the secound no is exclusive 

so meaning that the `a` will be included in the slice where the `b` will not be included

Defaults:

```go
s[:b]   // from 0 to b
s[a:]   // from a to len
s[:]    // whole slice
```

Capacity rule:
Capacity = capacity of original slice − start index

---

# Append (Dynamic Growth)


The built-in append function is used to dynamically add elements to a slice:

```go
func append(slice []Type, elems ...Type) []Type
```


```go
s = append(s, 10)
```

If capacity is exceeded:

* Go allocates a new larger array
* Copies data
* Returns new slice

Example:

```go
s := []int{1,2,3}
s = append(s, 4)
```

⚠ Always assign back:

```go
s = append(s, value)
```

If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

Notice that append() is variadic, the following are all valid:

>[!NOTE]
> At the last of this page i have talked little about varidic functions

```go 
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
```

---

# Copying Slices

Use `copy()`:

```go
a := []int{1,2,3}
b := make([]int, len(a))
copy(b, a)
```

This performs a **deep copy of elements** (but not nested structures).

---

# Nil vs Empty Slice

```go
var s []int        // nil slice
s := []int{}       // empty slice
```

Differences:

| Property | nil slice | empty slice |
| -------- | --------- | ----------- |
| len      | 0         | 0           |
| cap      | 0         | 0           |
| == nil   | true      | false       |

Both behave similarly in most cases.

---

# How Slice Sharing Works

Slices share the same underlying array.

Example:

```go
arr := []int{1,2,3,4}
a := arr[0:2]
b := arr[1:3]

b[0] = 100
fmt.Println(arr)
```

Output:

```
[1 100 3 4]
```

Why? Because `a`, `b`, and `arr` reference the same underlying array.

---

# Capacity Growth Strategy

When growing:

* Small slices → capacity roughly doubles
* Large slices → grows by smaller factor (~1.25x)

This makes append amortized O(1).

---

# Slices as Function Parameters

Slices are passed by value, but:

* The slice header is copied
* Underlying array is shared

So modifying elements inside function affects original slice.

But re-allocating inside function does NOT affect caller unless returned.

Example:

```go
func modify(s []int) {
    s[0] = 100
}
```

This changes original data.

---

# Reslicing

You can reslice within capacity:

```go
s := []int{1,2,3,4,5}
a := s[:3]
b := a[:5]   // valid if within capacity
```

---

# Memory Leak Gotcha

Large underlying array may stay in memory:

```go
big := make([]byte, 1_000_000)
small := big[:10]
```

Even though `small` has len=10, entire 1MB array stays in memory.

Fix:

```go
smallCopy := make([]byte, 10)
copy(smallCopy, big[:10])
```

---

# Multi-Dimensional Slices

```go
matrix := [][]int{
    {1,2,3},
    {4,5,6},
}
```

Each row is a separate slice.


# Slice of Slices

Slices can hold other slices, effectively creating a matrix, or a 2D slice.

```go 
rows := [][]int{}
rows = append(rows, []int{1, 2, 3})
rows = append(rows, []int{4, 5, 6})
fmt.Println(rows)
// [[1 2 3] [4 5 6]]
```

---

# Common Slice Patterns

**Remove element at index i**

```go
s = append(s[:i], s[i+1:]...)
```

**Insert at index i**

```go
s = append(s[:i], append([]T{value}, s[i:]...)...)
```

**Reverse slice**

```go
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
}
```

---

# Performance Characteristics

| Operation            | Complexity     |
| -------------------- | -------------- |
| Index access         | O(1)           |
| Append               | Amortized O(1) |
| Insert/Delete middle | O(n)           |
| Copy                 | O(n)           |

---

# When to Use Slices

Use slices:

* For dynamic collections
* When size not known at compile time
* For almost all list-like behavior in Go

Use arrays:

* Rarely
* When fixed size is required
* When used internally for performance guarantees

---


# Variadic Function

Functions can take an arbitrary number of final arguments. This is done using the ... syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

```go 
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
```

The familiar fmt.Println() and fmt.Sprintf() are variadic, as are many in the standard library! fmt.Println() prints each element with space delimiters and a newline at the end.


```go 
func Println(a ...interface{}) (n int, err error)
```


**Spread Operator**

The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

```go
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
```

---

<br>
<br>
<br>
<br>
<br>



>[!WARNING]
> plz see the Tricky Slices


# Tricky Slices

## 🔹 Key Points

1. **Slices have length and capacity**

   * `len(slice)` = number of elements you can access
   * `cap(slice)` = size of the underlying array

2. **append() may or may not allocate a new array**

   * If the slice has **enough capacity**, `append()` reuses the existing array.
   * If capacity is exceeded, `append()` allocates a **new array**.

3. **Result of append() should be assigned to a slice variable**

   * Always do: `s = append(s, x)`
   * Avoid doing: `otherSlice = append(slice, x)` unless you’re intentionally merging slices

4. **Shared underlying array can cause unexpected overwrites**

   * If two slices share the same array, appending to one can modify elements seen by the other slice if capacity allows.

---

## 🔹 Simple Example

```go
i := make([]int, 3, 5)  // len=3, cap=5
i[0], i[1], i[2] = 1, 2, 3

j := append(i, 4)       // uses same underlying array
g := append(i, 5)       // also uses same underlying array

fmt.Println("i:", i)    // [1 2 3]
fmt.Println("j:", j)    // [1 2 3 5] <- g overwrote j's 4
fmt.Println("g:", g)    // [1 2 3 5]
```

✅ Explanation:

* `i` has capacity 5 → enough for appending without allocating a new array.
* `j` and `g` **share the same underlying array**.
* Appending to `g` can overwrite `j`’s data.

---

## 🔹 Safe Pattern

Always assign append result to the slice you’re modifying:

```go
s := []int{1, 2, 3}
s = append(s, 4)
fmt.Println(s) // [1 2 3 4]
```

* This avoids surprises with shared underlying arrays.

---


