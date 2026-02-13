## `fmt.Fprintf()` in Go

`Fprintf` stands for **File Print Formatted** (more generally: **Formatted Print to Writer**).

It formats text and writes it to an `io.Writer`.

---

## Basic Syntax

```go
fmt.Fprintf(writer, format, values...)
```

* `writer` → any value that implements the `io.Writer` interface
* `format` → format string (`%s`, `%d`, etc.)
* `values...` → variables to insert

It returns:

```go
n int, err error
```

* `n` → number of bytes written
* `err` → error (if any)

---

## Writing to Standard Output

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Bilal"
	fmt.Fprintf(os.Stdout, "Hello %s\n", name)
}
```

This behaves like `Printf`, because `os.Stdout` is the default output.

---

## Writing to a File

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Age: %d\n", 25)
}
```

Now the formatted text goes into `output.txt`.

---

## Writing to a Buffer (In-Memory)

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Pi: %.2f", 3.14159)

	fmt.Println(buf.String())
}
```

Useful for:

* Building dynamic strings
* Logging systems
* HTTP responses

---

## Format Verbs

`Fprintf` uses the same format verbs as `Printf` and `Sprintf`:

* `%s` → string
* `%d` → integer
* `%f` → float
* `%.2f` → float precision
* `%t` → bool
* `%v` → default format
* `%T` → type
* `%q` → quoted string

---

## Difference Between Printf, Sprintf, and Fprintf

| Function  | Output Destination | Returns String? |
| --------- | ------------------ | --------------- |
| `Printf`  | Standard output    | ❌ No            |
| `Sprintf` | Returns string     | ✅ Yes           |
| `Fprintf` | Any `io.Writer`    | ❌ No            |

---
