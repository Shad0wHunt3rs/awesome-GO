## `Printf` example for **each kind of variable**.

* Printf is a function from the fmt package in Go.
* Its full name is “Print Formatted”.
* It allows you to print variables and text in a formatted way, similar to printf in C.
* The f in Printf stands for “formatted”.

| Specifier | Description                                                       | Example    |
| --------- | ----------------------------------------------------------------- | ---------- |
| `%d`      | Decimal integer                                                   | 42         |
| `%f`      | Floating-point                                                    | 3.14       |
| `%.2f`    | Floating-point with 2 digits                                      | 3.14       |
| `%t`      | Boolean                                                           | true       |
| `%s`      | String                                                            | "Hello"    |
| `%c`      | Character / rune                                                  | 'A'        |
| `%v`      | Default format (works for complex numbers, structs, slices, etc.) | 3+4i       |
| `%#v`     | Go-syntax representation of the value                             | int32(-42) |

---

### **1. Boolean**

```go
package main

import "fmt"

func main() {
    flag := true
    fmt.Printf("Boolean flag: %t\n", flag) // %t for booleans
}
```

---

### **2. String**

```go
package main

import "fmt"

func main() {
    name := "Alice"
    fmt.Printf("String name: %s\n", name) // %s for strings
}
```

---

### **3. Signed Integers**

```go
package main

import "fmt"

func main() {
    i := int(-42)
    i8 := int8(-8)
    i16 := int16(-32000)
    i32 := int32(-2000000000)
    i64 := int64(-9000000000000000000)

    fmt.Printf("Signed integers: i=%d, i8=%d, i16=%d, i32=%d, i64=%d\n", i, i8, i16, i32, i64)
}
```

---

### **4. Unsigned Integers**

```go
package main

import "fmt"

func main() {
    u := uint(42)
    u8 := uint8(255)
    u16 := uint16(60000)
    u32 := uint32(4000000000)
    u64 := uint64(18000000000000000000)
    uptr := uintptr(123456789)

    fmt.Printf("Unsigned integers: u=%d, u8=%d, u16=%d, u32=%d, u64=%d, uintptr=%d\n", u, u8, u16, u32, u64, uptr)
}
```

---

### **5. Aliases (`byte` and `rune`)**

```go
package main

import "fmt"

func main() {
    b := byte(100)        // alias for uint8
    r := rune('♥')        // alias for int32, Unicode character

    fmt.Printf("byte value: %d\n", b)   // %d for integer
    fmt.Printf("rune value: %c\n", r)   // %c for character
}
```

---

### **6. Floating-Point Numbers**

```go
package main

import "fmt"

func main() {
    f32 := float32(3.14)
    f64 := float64(3.1415926535)

    fmt.Printf("Float32: %f\n", f32)
    fmt.Printf("Float64 with 10 digits: %.10f\n", f64)
}
```

---

### **7. Complex Numbers**

```go
package main

import "fmt"

func main() {
    c64 := complex64(complex(3, 4))
    c128 := complex128(complex(5.5, -2.2))

    fmt.Printf("Complex64: %v\n", c64)            // %v prints complex number
    fmt.Printf("Complex128: %v\n", c128)
    fmt.Printf("Complex128 parts: real=%.2f imag=%.2f\n", real(c128), imag(c128)) // separate parts
}
```

---

Here are additional important fmt format specifiers in Go:

| Specifier | Description                                 | Example             |
| --------- | ------------------------------------------- | ------------------- |
| `%b`      | Binary representation (int)                 | 101010              |
| `%o`      | Octal representation                        | 52                  |
| `%O`      | Octal with `0o` prefix                      | 0o52                |
| `%x`      | Lowercase hexadecimal                       | 2a                  |
| `%X`      | Uppercase hexadecimal                       | 2A                  |
| `%e`      | Scientific notation (lowercase)             | 3.140000e+00        |
| `%E`      | Scientific notation (uppercase)             | 3.140000E+00        |
| `%g`      | Compact float format (chooses `%e` or `%f`) | 3.14                |
| `%G`      | Compact float (uppercase)                   | 3.14                |
| `%p`      | Pointer address                             | 0xc0000140a0        |
| `%q`      | Quoted string                               | `"Hello"`           |
| `%+v`     | Struct with field names                     | `{Name:Ali Age:20}` |
| `%T`      | Type of the value                           | int                 |
| `%%`      | Prints a literal `%` sign                   | %                   |
| `%U`      | Unicode format (`U+XXXX`)                   | U+0041              |


