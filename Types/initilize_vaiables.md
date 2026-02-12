Go has several ways to initialize variables, depending on the scope, type, and whether you want to specify the type explicitly or let Go infer it.


## **1. Using `var` with explicit type**

```go
package main

import "fmt"

func main() {
    var age int = 25         // Declare a variable 'age' of type int and initialize with 25
    var name string = "Alice" // Declare a string variable 'name' and initialize

    fmt.Println("Age:", age)
    fmt.Println("Name:", name)
}
```

* Good for clarity.
* The type must match the value.

---

## **2. Using `var` without explicit type (type inference)**

```go
package main

import "fmt"

func main() {
    var age = 25       // Go infers 'int' from the value
    var name = "Bob"   // Go infers 'string' from the value

    fmt.Println(age, name)
}
```

* Go figures out the type automatically.
* Saves typing and is very common.

---

## **3. Short variable declaration (`:=`)**

```go
package main

import "fmt"

func main() {
    age := 30          // Declare and initialize in one step, type inferred as int
    name := "Charlie"  // Type inferred as string

    fmt.Println(age, name)
}
```

* Only works **inside functions**, not at package level.
* Most common way in Go for local variables.

---

## **4. Declaring multiple variables at once**

```go
package main

import "fmt"

func main() {
    var x, y, z int = 1, 2, 3    // Multiple variables with explicit type
    a, b, c := "Go", 3.14, true   // Multiple variables with type inference

    fmt.Println(x, y, z)
    fmt.Println(a, b, c)
}
```

---

## **5. Zero-value initialization**

```go
package main

import "fmt"

func main() {
    var number int       // Not initialized explicitly
    var flag bool        // Not initialized explicitly
    var text string      // Not initialized explicitly

    fmt.Println(number) // 0 (default int value)
    fmt.Println(flag)   // false (default bool value)
    fmt.Println(text)   // "" (empty string)
}
```

* Every variable has a **default “zero value”** if not initialized.
* Examples:

  * `int` → `0`
  * `float64` → `0.0`
  * `bool` → `false`
  * `string` → `""`

---

### **Summary Table**

| Method             | Example        | Type Inferred? | Scope         |
| ------------------ | -------------- | -------------- | ------------- |
| `var x int = 10`   | var x int = 10 | No             | Package/func  |
| `var x = 10`       | var x = 10     | Yes            | Package/func  |
| `x := 10`          | x := 10        | Yes            | Function only |
| Multiple variables | var a,b=1,2    | Can infer      | Both          |


> [!NOTE] 
> Type inference means Go automatically figures out the variable’s type based on the value you assign to it.
> You don’t have to explicitly declare the type if Go can guess it from the value.
