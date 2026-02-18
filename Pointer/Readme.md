# **Pointers in Go**

### **What is a Pointer?**

A **pointer** is a variable that **stores the memory address** of another variable, rather than its value.

* Think of it like a “map” to the actual data in memory.
* The type of a pointer in Go is written as `*T`, where `T` is the type of the value it points to.

```go
var x int = 10
var p *int = &x // p points to x
```

* `&x` → address-of operator, gives the memory address of `x`
* `*p` → dereference operator, gives the value at the address `p` points to

```go
fmt.Println(p)  // prints memory address of x
fmt.Println(*p) // prints 10
```

---

### **Zero Value of a Pointer**

* A pointer that hasn’t been assigned **defaults to `nil`**.
* Accessing a nil pointer causes a **runtime panic**.

```go
var p *int
fmt.Println(p) // <nil>
```

---

### **Basic Pointer Operations**

| Operation          | Example   | Meaning                           |
| ------------------ | --------- | --------------------------------- |
| Get address        | `&x`      | Address of variable               |
| Dereference        | `*p`      | Value at pointer                  |
| Assign via pointer | `*p = 20` | Change original value via pointer |

**Example:**

```go
x := 5
p := &x
fmt.Println(*p) // 5

*p = 10        // modify x through pointer
fmt.Println(x) // 10
```

---

### **Pointers vs Values**

* Assigning a value **copies it**, but assigning a pointer **copies the address**, not the data.

```go
a := 10
b := a     // copy of value
b = 20
fmt.Println(a) // 10

pa := &a
pb := pa    // copy of pointer (both point to same address)
*pb = 30
fmt.Println(a) // 30
```

---

### **Pointers and Functions**

#### **Passing by Value (normal)**

* Go passes arguments **by value** (copy). Changes inside function **don’t affect the caller**.

```go
func increment(x int) {
    x++
}

a := 5
increment(a)
fmt.Println(a) // 5 (unchanged)
```

#### **Passing by Pointer**

* Pass the **address** to allow function to modify original variable.

```go
func incrementPtr(x *int) {
    *x++
}

a := 5
incrementPtr(&a)
fmt.Println(a) // 6
```

This is how you **mutate variables inside functions** in Go.

---

### **Pointer to Pointer**

* Go allows pointers to pointers (rarely used):

```go
x := 10
p := &x       // *int
pp := &p      // **int
fmt.Println(**pp) // 10
```

---

### **Pointers and Structs**

Pointers are very useful for **structs**, especially large ones:

```go
type Point struct{ X, Y int }
p := Point{1, 2}

pp := &p        // pointer to struct
pp.X = 10       // shorthand: Go automatically dereferences
fmt.Println(p)  // {10, 2}
```

> Note: Go lets you access struct fields via pointer without explicitly writing `(*pp).X`.

---

### **Pointers and Arrays / Slices**

#### **Arrays**

* Arrays are **value types**, so passing an array copies it.
* Using pointers avoids copy:

```go
arr := [3]int{1, 2, 3}
p := &arr
(*p)[0] = 10
fmt.Println(arr) // [10 2 3]
```

#### **Slices**

* Slices are **already reference types**.
* Passing a slice to a function does **not require pointers** to modify elements:

```go
func modifySlice(s []int) {
    s[0] = 99
}

arr := []int{1,2,3}
modifySlice(arr)
fmt.Println(arr) // [99 2 3]
```

> Only use pointers with slices if you want to **change the slice itself (length, capacity, or underlying array)**.

---

### **Pointers and Maps**

* Maps in Go are **reference types**.
* You rarely need pointers to maps to modify their contents:

```go
m := map[string]int{"a":1}
func addKey(mp map[string]int) {
    mp["b"] = 2
}
addKey(m)
fmt.Println(m) // map[a:1 b:2]
```

* Use a pointer to a map **only if you want to reassign the map itself** inside a function.

---

### **Pointers and nil**

* Only dereference if pointer is **non-nil**:

```go
var p *int
if p != nil {
    fmt.Println(*p)
} else {
    fmt.Println("pointer is nil")
}
```

* Prevents **runtime panics**.

---

### **Why Pointers in Go Are Useful**

1. **Avoid copying large structs or arrays**
2. **Modify variables inside functions**
3. **Efficient memory usage**
4. **Linked data structures** like linked lists, trees
5. **Optional values** (nil pointers can indicate “no value”)

---

### **Restrictions / Notes**

* No pointer arithmetic in Go (unlike C/C++)
* Only take address of **addressable variables** (e.g., local variables, struct fields, array elements)
* Cannot take address of **constants or function return values**

```go
x := 10
ptr := &x   // OK
ptr2 := &5  // ❌ illegal
```

---
