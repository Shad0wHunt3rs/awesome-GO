# Packages in Go


A **package** is a collection of Go source files in the same directory that are compiled together.

It is:

* A **namespace**
* A **code organization unit**
* A **visibility boundary**

Every Go file must start with:

```go
package packagename
```

---

## The `main` Package

```go
package main
```

* Special package.
* Required for executable programs.
* Must contain:

```go
func main()
```

When you run:

```bash
go run main.go
```

Go builds the `main` package into a binary and executes `main()`.

If a package is NOT `main`, it becomes a **library package**.

---

## Custom Packages

Example project:

```
project/
│
├── main.go
└── mathutils/
    └── add.go
```

### `mathutils/add.go`

```go
package mathutils

func Add(a, b int) int {
    return a + b
}
```

### `main.go`

```go
package main

import (
    "fmt"
    "project/mathutils"
)

func main() {
    fmt.Println(mathutils.Add(2, 3))
}
```

---

## Exported vs Unexported Identifiers

This is VERY IMPORTANT.

In Go:

| Capitalized | Exported |
| ----------- | -------- |
| lowercase   | Private  |

Example:

```go
func Add() {}     // exported
func add() {}     // not exported
```

Go does not use `public` or `private`.
Capitalization controls visibility.

This is enforced by the compiler.

---

## Package Scope

Each package has its own namespace.

Two packages can have:

```go
func Init()
```

No conflict unless imported with same name.

---

## Importing Packages

### Standard Library

```go
import "fmt"
import "math"
```

Example:

```go
fmt.Println(math.Sqrt(16))
```

---

### Multiple Imports

```go
import (
    "fmt"
    "math"
)
```

---

### Alias Import

```go
import m "math"

m.Sqrt(4)
```

---

### Blank Import (Side Effects)

```go
import _ "some/package"
```

Used when:

* You want `init()` to run
* But don't need to reference exported identifiers

Example:
Used in database drivers.

---

### Dot Import (NOT recommended)

```go
import . "math"
```

Now you can call:

```go
Sqrt(4)
```

This pollutes namespace. Avoid it.

---

## `init()` Function

Every package can have:

```go
func init()
```

* Runs automatically
* Runs before `main()`
* Used for setup

Order:

1. Imported packages
2. Current package
3. main()

If multiple files → init runs in file order.

---

## Package Initialization Order (Important)

Go builds a dependency graph.

If:

* A imports B
* B imports C

Initialization order:

C → B → A → main

No circular imports allowed.

---

>[!NOTE]
> In Go, each folder corresponds to a single package, so all `.go` files in the same directory must declare the same `package` name. Mixing package names in one folder will cause a compiler error. For example, if you have `add.go` and `mul.go`, both must say `package mathutils`. You cannot put `package helper` in the same folder. Go treats the folder as one compilation unit, which keeps builds simple and fast. If you need another package, create a separate folder. The only exception is test files: a file ending in `_test.go` can use `packagename_test` to test exported functions. The key rule to remember is: **Directory = Package**.

example:

```bash
/mathutils
    add.go
    sub.go
```

Both: 

```bash
package mathutils
```



# Modules in Go

Now we go deeper.

A **module** is a collection of related packages.

It defines:

* Dependency management
* Versioning
* Project boundary

Before modules → GOPATH system.
Now → Go Modules (modern system).

---

## Creating a Module

Inside your project folder:

```bash
go mod init github.com/username/projectname
```

This creates:

```
go.mod
```

Example:

```go
module github.com/muhammadbilal/myapp

go 1.22
```

This defines:

* Module path
* Go version

---

## Why Module Path Matters

When importing:

```go
import "github.com/muhammadbilal/myapp/mathutils"
```

The import path starts with module name.

Even locally, Go uses module path.

---

## `go.mod` File

Example:

```go
module github.com/muhammadbilal/myapp

go 1.22

require github.com/gorilla/mux v1.8.0
```

It contains:

* Module name
* Go version
* Dependencies
* Dependency versions

---

## `go.sum`

Automatically generated.

Contains:

* Cryptographic hashes
* Ensures dependency integrity
* Prevents tampering

Never manually edit it.

---

## Adding Dependencies

If you write:

```go
import "github.com/gorilla/mux"
```

Then run:

```bash
go mod tidy
```

Go:

* Downloads dependency
* Adds to go.mod
* Updates go.sum

---

## `go mod tidy`

Very important command.

It:

* Removes unused dependencies
* Adds missing ones
* Cleans go.mod

Always run before pushing code.

---

## Versioning System

Go uses Semantic Versioning:

```
vMAJOR.MINOR.PATCH
```

Example:

```
v1.2.3
```

* MAJOR → breaking change
* MINOR → feature
* PATCH → bug fix

---

## Module Proxy & Caching

Go downloads modules into:

```
$GOPATH/pkg/mod
```

It uses:

* Proxy server
* Module cache
* Checksums

So builds are reproducible.

---

# Internal Packages

If you create:

```
project/
    internal/
        helper.go
```

Packages inside `internal/`:

* Cannot be imported outside parent module

This enforces private module-level boundaries.

Used in large systems.

---
