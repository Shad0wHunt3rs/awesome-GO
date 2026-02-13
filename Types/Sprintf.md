## `fmt.Sprintf()` in Go

`Sprintf` stands for **String Print Formatted**.

It formats values according to a format specifier and **returns the result as a string** instead of printing it.

---

## Basic Syntax

```go
result := fmt.Sprintf(format, values...)
```

* `format` → string with format verbs (`%s`, `%d`, etc.)
* `values...` → variables to insert
* Returns a **string**

---

## Simple Example

```go
package main

import "fmt"

func main() {
	name := "Bilal"
	age := 20

	result := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(result)
}
```

Output:

```
Name: Bilal, Age: 20
```

---

## Why Use `Sprintf`?

Unlike `Printf`, it **does not print immediately**.

It lets you:

* Store formatted strings
* Return formatted messages from functions
* Build dynamic strings
* Log messages
* Create JSON-like or structured text

---

## 4️⃣ Formatting Verbs (Commonly Used)

| Verb   | Meaning                  | Example   |
| ------ | ------------------------ | --------- |
| `%s`   | string                   | "Hello"   |
| `%d`   | integer                  | 42        |
| `%f`   | float                    | 3.14159   |
| `%.2f` | float (2 decimal places) | 3.14      |
| `%t`   | boolean                  | true      |
| `%v`   | default format           | any type  |
| `%T`   | type of value            | int       |
| `%q`   | quoted string            | `"Hello"` |
| `%x`   | hex                      | 2a        |
| `%p`   | pointer address          | 0xc00001  |

---

## Controlling Width and Alignment

### Minimum Width

```go
fmt.Sprintf("%10s", "Go")
```

Output:

```
        Go
```

Right-aligned in 10 spaces.

---

### Left Alignment

```go
fmt.Sprintf("%-10s", "Go")
```

Output:

```
Go        
```

---

### Zero Padding

```go
fmt.Sprintf("%05d", 42)
```

Output:

```
00042
```

---

## Formatting Floats

```go
price := 99.45678
fmt.Sprintf("%.2f", price)
```

Output:

```
99.46
```

### Scientific Notation

```go
fmt.Sprintf("%e", 1234.5)
```

Output:

```
1.234500e+03
```

---

## Struct Formatting

```go
type User struct {
	Name string
	Age  int
}

u := User{"Ali", 21}

fmt.Sprintf("%v", u)
```

Output:

```
{Ali 21}
```

With field names:

```go
fmt.Sprintf("%+v", u)
```

Output:

```
{Name:Ali Age:21}
```

Go-syntax representation:

```go
fmt.Sprintf("%#v", u)
```

---


## Difference

| Function  | Prints?          | Returns string? |
| --------- | ---------------- | --------------- |
| `Printf`  | ✅ Yes            | ❌ No            |
| `Sprintf` | ❌ No             | ✅ Yes           |
| `Fprintf` | Writes to writer | ❌ No            |

---
