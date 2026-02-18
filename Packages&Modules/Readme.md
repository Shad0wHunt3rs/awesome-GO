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

## 1.1 The `main` Package

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

## 1.2 Custom Packages

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

## 1.3 Exported vs Unexported Identifiers

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

## 1.4 Package Scope

Each package has its own namespace.

Two packages can have:

```go
func Init()
```

No conflict unless imported with same name.

---

## 1.5 Importing Packages

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

## 1.6 `init()` Function

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

## 1.7 Package Initialization Order (Important)

Go builds a dependency graph.

If:

* A imports B
* B imports C

Initialization order:

C → B → A → main

No circular imports allowed.

---

# 2️⃣ Modules in Go

Now we go deeper.

## What is a Module?

A **module** is a collection of related packages.

It defines:

* Dependency management
* Versioning
* Project boundary

Before modules → GOPATH system.
Now → Go Modules (modern system).

---

## 2.1 Creating a Module

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

## 2.2 Why Module Path Matters

When importing:

```go
import "github.com/muhammadbilal/myapp/mathutils"
```

The import path starts with module name.

Even locally, Go uses module path.

---

## 2.3 `go.mod` File

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

## 2.4 `go.sum`

Automatically generated.

Contains:

* Cryptographic hashes
* Ensures dependency integrity
* Prevents tampering

Never manually edit it.

---

## 2.5 Adding Dependencies

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

## 2.6 `go mod tidy`

Very important command.

It:

* Removes unused dependencies
* Adds missing ones
* Cleans go.mod

Always run before pushing code.

---

## 2.7 Versioning System

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

## 2.8 Module Proxy & Caching

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

# 3️⃣ Internal Packages

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

# 4️⃣ Difference: Package vs Module

| Package                 | Module                 |
| ----------------------- | ---------------------- |
| Collection of .go files | Collection of packages |
| Directory level         | Project level          |
| Namespace unit          | Dependency unit        |
| Controls visibility     | Controls versioning    |

---

# 5️⃣ Advanced Concepts

## 5.1 Replace Directive

In `go.mod`:

```go
replace github.com/old/package => ../local/package
```

Used for:

* Local development
* Testing modified dependencies

---

## 5.2 Require Directive

```go
require github.com/gorilla/mux v1.8.0
```

Defines dependency version.

---

## 5.3 Indirect Dependencies

```go
require github.com/pkg/errors v0.9.1 // indirect
```

Means:
Dependency of your dependency.

---

## 5.4 Multi-Module Workspace (Go 1.18+)

Use:

```bash
go work init
```

Creates:

```
go.work
```

Allows multiple modules in one workspace.

Used in monorepos.

---

# 6️⃣ How Go Compiler Handles Packages

When building:

1. Parses imports
2. Builds dependency graph
3. Compiles bottom-up
4. Caches compiled packages
5. Links into binary

This is why Go builds are fast.

---

# 7️⃣ Real Project Structure (Professional)

```
myapp/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── auth/
│   └── database/
│
├── pkg/
│   └── utils/
│
├── go.mod
└── go.sum
```

### Explanation:

* `cmd/` → entry points
* `internal/` → private app logic
* `pkg/` → reusable libraries
* `go.mod` → module root

---

# 8️⃣ Common Beginner Mistakes

1. Forgetting `go mod init`
2. Using relative imports (not allowed)
3. Circular imports
4. Not capitalizing exported functions
5. Manually editing go.sum

---

# 9️⃣ Important Commands Summary

```bash
go mod init
go mod tidy
go mod download
go mod vendor
go list
go build
go install
```

---

# 🔟 Mental Model (Very Important)

Think like this:

* **File** → contains code
* **Package** → group of files (namespace)
* **Module** → group of packages (dependency unit)
* **Workspace** → group of modules

Hierarchy:

```
Workspace
  └── Module
        └── Package
              └── File
```

---
