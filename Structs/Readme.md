## Structs in Go

In **Go**, a `struct` (short for *structure*) is a composite data type used to group variables together under one name. It is similar to a class in other languages, but without inheritance.

Structs are fundamental in Go because Go is built around **composition**, not classical OOP inheritance.

---

# Basic Struct Definition

**Syntax**

```go
type StructName struct {
    fieldName fieldType
    fieldName2 fieldType2
}
```

Example

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

Here:

* `Person` is a new type
* It contains three fields: `Name`, `Age`, and `City`

---

# Creating Struct Instances

### Method 1: Zero Value Initialization

```go
var p Person
```

All fields get default (zero) values:

* string → ""
* int → 0
* bool → false
* pointer → nil

---

### Method 2: Struct Literal (Recommended)

```go
p := Person{
    Name: "Ali",
    Age:  21,
    City: "Lahore",
}
```

This is the safest and most readable way.

---

### Method 3: Positional Initialization (Not Recommended)

```go
p := Person{"Ali", 21, "Lahore"}
```

⚠ Risky because order must match exactly.

---

# Accessing Fields

```go
fmt.Println(p.Name)
p.Age = 25
```

Use dot `.` operator.

---

# Structs and Memory

Structs in Go are **value types**.

This means:

```go
p1 := Person{"Ali", 20, "Karachi"}
p2 := p1

p2.Name = "Ahmed"

fmt.Println(p1.Name) // Still "Ali"
```

Because a **copy** is made.

---

# 5. Using Pointers to Structs

To avoid copying:

```go
p := &Person{
    Name: "Ali",
    Age:  20,
}
```

Accessing fields:

```go
fmt.Println(p.Name) // Go automatically dereferences
```

Internally this is:

```go
(*p).Name
```

---

# Structs as Function Parameters

### Pass by Value

```go
func updateAge(p Person) {
    p.Age = 30
}
```

Changes won't affect original.

---

### Pass by Pointer

```go
func updateAge(p *Person) {
    p.Age = 30
}
```

Now original is modified.

---

# Methods on Structs

Go attaches methods to types.

### Value Receiver

```go
func (p Person) greet() string {
    return "Hello " + p.Name
}
```

Because the function has a receiver, it becomes a method attached to the Person type. Without a receiver, it would just be a normal function.

Called like:

```go
p.greet()
```

Is actually syntactic sugar for:

```go
greet(p)
```

Go just makes it look like OOP

---

### Pointer Receiver

```go
func (p *Person) haveBirthday() {
    p.Age++
}
```

### Rule:

* Use pointer receiver if method modifies struct
* Use pointer receiver if struct is large (avoid copying)

---

# Anonymous Structs

Structs without type name.

```go
p := struct {
    Name string
    Age  int
}{
    Name: "Ali",
    Age:  20,
}
```

* Because sometimes you need a struct only once.
* You don’t want to create a new type just for one small use.

Useful for:

* Temporary data
* JSON responses
* Tests

---

# Nested Structs

```go
type Address struct {
    City    string
    ZipCode string
}

type Person struct {
    Name    string
    Age     int
    Address Address
}
```

Access:

```go
p.Address.City
```

you can assign values to a nested struct in Go.


```go
p := Person{
    Name: "Ali",           // Assign Name field
    Age:  20,              // Assign Age field
    Address: Address{       // Assign Address field, which is itself a struct
        City:    "Karachi",
        ZipCode: "74400",
    },
}
```


---

# Embedded Structs (Composition)

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}
```

Now:

```go
d := Dog{
    Animal: Animal{Name: "Tommy"},
    Breed:  "Bulldog",
}

fmt.Println(d.Name) // promoted field
```

This is called **field promotion**.

Go uses composition instead of inheritance.


| Feature             | Nested Struct          | Embedded Struct                 |
| ------------------- | ---------------------- | ------------------------------- |
| Declaration         | `FieldName Type`       | `Type` (no field name)          |
| Access Fields       | `p.FieldName.SubField` | `p.SubField` (promoted)         |
| Methods             | `p.Field.Method()`     | `p.Method()` (promoted)         |
| Reuse / Composition | Normal field           | Field + promotion → composition |
| Common Use          | Data grouping          | Composition / mix-in            |

* Nested struct → like a box inside a box. You always go through the box.
* Embedded struct → the box’s contents are poured directly into the outer box. You can access them directly.

Use **nested struct** when you want a **clear, separate field** and don’t need to access inner fields directly.
Use **embedded struct** when you want **field and method promotion** for **composition or “mix-in” style reuse** — it’s convenient but can be confusing if overused.


---

# Struct Tags

Used for metadata (especially JSON, DB, validation).

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

Used by packages like:

* encoding/json

Example:

```go
json.Marshal(user)
```

Here’s a concise list of Go libraries/packages that commonly use struct tags:

1. `encoding/json`
2. `encoding/xml`
3. `encoding/gob`
4. `github.com/json-iterator/go`
5. `database/sql`
6. `gorm.io/gorm`
7. `github.com/jmoiron/sqlx`
8. `upper.io/db.v3`
9. `github.com/go-playground/validator/v10`
10. `github.com/asaskevich/govalidator`
11. `github.com/mitchellh/mapstructure`
12. `github.com/mitchellh/hashstructure`
13. `github.com/golang/protobuf`
14. `github.com/go-playground/form`
15. `github.com/robfig/cron/v3`

These are the main ones where struct tags are used for JSON, DB mapping, validation, encoding, or other metadata purposes.

---

# Comparing Structs

Structs are comparable **only if all fields are comparable**.

Allowed:

* int
* string
* bool
* arrays

Not allowed:

* slices
* maps
* functions

```go
p1 == p2
```

---

# Structs with Slices and Maps

```go
type Student struct {
    Name   string
    Marks  []int
    Grades map[string]string
}
```

Important:

* Slice and map inside struct are references
* But struct itself is still copied

---

# Constructor Pattern (Idiomatic Go)

Go has no constructors, but convention is:

```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
```

Usage:

```go
p := NewPerson("Ali", 20)
```

---

# Exported vs Unexported Fields

Capital letter = exported

```go
type Person struct {
    Name string  // exported
    age  int     // private
}
```

If used in another package:

* `Name` accessible
* `age` not accessible

---

# Empty Struct

```go
struct{}
```

Takes **zero bytes** of memory.

Used for:

* Sets (map[string]struct{})
* Signaling channels

---

# Memory Layout

Go arranges struct fields in memory with padding for alignment.


the order of fields in a struct can have a big impact on memory usage. This is the same struct as above, but poorly designed:

Example:

```go
type stats struct {
    NumPosts uint8
	Reach    uint16
	NumLikes uint8
}
```

```bash
Size of stats struct: 6 bytes
```

Notice that Go has "aligned" the fields, meaning that it has added some padding (wasted space) to make up for the size difference between the uint16 and uint8 types. It's done for execution speed, but it can lead to increased memory usage.


Optimization:

```go
type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}
```

```bash
Size of stats struct: 4 bytes
```


Better memory alignment.


* Order struct fields from largest to smallest to reduce padding and improve memory alignment.
* Group small fields together to improve cache efficiency and speed, even if total size stays the same.

---

# Struct vs Interface

Struct:

* Concrete data

Interface:

* Behavior definition

Example:

```go
type Shape interface {
    Area() float64
}
```

Struct implements interface implicitly.

---

# Structs and JSON Example

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

```go
data, _ := json.Marshal(user)
```

Unmarshal:

```go
var u User
json.Unmarshal(data, &u)
```

---

# When to Use Structs

Use structs when:

* Modeling real-world entities
* Grouping related data
* Building APIs
* Designing systems
* Replacing classes in OOP languages

---

# Key Takeaways

* Structs are value types.
* Use pointers to avoid copying.
* Use methods for behavior.
* Go prefers composition over inheritance.
* Struct tags add metadata.
* Empty struct uses zero memory.
* Interfaces are implemented implicitly.

---

