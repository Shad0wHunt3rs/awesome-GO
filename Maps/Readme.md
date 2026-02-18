# Maps in Go

In Go, a **map** is a built-in reference type used to associate **keys** with **values** (like dictionaries in Python or hash maps in C++/Java).

Internally, Go maps are implemented as **hash tables**.

---

## Basic Syntax

**Declaration**

```go
var m map[string]int
```

This declares a map but **does not initialize it**.
The zero value of a map is `nil`.

If you try to insert into it:

```go
m["a"] = 10  // panic: assignment to entry in nil map
```

It will panic.

---

## Creating (Initializing) a Map

**Using `make` (most common)**

```go
m := make(map[string]int)
```

With capacity hint:

```go
m := make(map[string]int, 100)
```

> The second argument is just a hint, not fixed size.

---

**Using map literal**

```go
m := map[string]int{
    "a": 1,
    "b": 2,
}
```

or we could also do is 

```go
m := map[string]int{}
```

---

**Adding and Updating Elements**

```go
m := make(map[string]int)

m["apple"] = 5   // add
m["apple"] = 10  // update
```

Maps automatically grow as needed.

---

**Accessing Values**

```go
value := m["apple"]
```

If the key does not exist:

* It returns **zero value** of the value type.
* No error.

Example:

```go
m := make(map[string]int)
fmt.Println(m["x"])  // 0
```

---

**Checking if Key Exists**

Because zero value is returned for missing keys, use **comma-ok idiom**:

```go
value, ok := m["apple"]
```

* `ok == true` → key exists
* `ok == false` → key does not exist

Example:

```go
if v, ok := m["x"]; ok {
    fmt.Println("Found:", v)
} else {
    fmt.Println("Not found")
}
```

---

**Deleting Elements**

```go
delete(m, "apple")
```

* Safe even if key doesn't exist.
* Does nothing if key is absent.

---

**Iterating Over a Map**

**Using range**

```go
for key, value := range m {
    fmt.Println(key, value)
}
```

**Important**:

Iteration order is **not guaranteed**.
Maps are deliberately randomized for security reasons.

If you need ordered iteration:

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
```

---

## Maps Are Reference Types

Maps are not copied when passed to functions.

```go
func modify(m map[string]int) {
    m["x"] = 100
}
```

Changes affect original map.

Example:

```go
m := make(map[string]int)
modify(m)
fmt.Println(m["x"])  // 100
```

---

## Comparing Maps

You **cannot** compare maps directly:

```go
if m1 == m2 { }  // compile error
```

Only allowed comparison:

```go
if m == nil
```

To compare maps, manually compare keys and values.

---

## Map with Struct Values

```go
type User struct {
    Name string
    Age  int
}

users := make(map[int]User)
users[1] = User{"Ali", 20}
```

---

## Map with Struct Keys

Allowed only if the struct is **comparable**.

Allowed key types:

* bool
* numeric
* string
* pointer
* channel
* interface (if underlying comparable)
* struct (if all fields comparable)
* array (if elements comparable)

Not allowed:

* slice
* map
* function

Example:

```go
type Point struct {
    X, Y int
}

m := make(map[Point]string)
m[Point{1, 2}] = "A"
```

---

## Map of Slices

Common pattern:

```go
m := make(map[string][]int)

m["a"] = append(m["a"], 1)
m["a"] = append(m["a"], 2)
```

This works because:

* Missing key → zero value of slice → nil slice
* `append` works on nil slices

---

## Nested Maps

```go
m := make(map[string]map[string]int)

m["group"] = make(map[string]int)
m["group"]["a"] = 1
```

Be careful:
You must initialize inner maps before use.

---

# Clearing a Map

Option 1:

```go
for k := range m {
    delete(m, k)
}
```

Option 2 (Go 1.21+):

```go
clear(m)
```

---

## Map Memory & Internal Working

### Go Maps Are Hash Tables

A Go `map[K]V` is implemented as a **hash table** under the hood.

* **Key** → hashed → **bucket index** → store value.
* The buckets contain entries for multiple key/value pairs (to handle collisions).
* Each bucket stores a small fixed number of entries (8 by default in Go 1.19+).

So conceptually:

```
map:
+--------+--------+--------+...
| bucket | bucket | bucket |
+--------+--------+--------+

bucket:
+--------+--------+--------+--------+--------+--------+--------+--------+
| key0   | val0   | key1   | val1   | ... 8 entries ...                 |
+--------+--------+--------+--------+--------+--------+--------+--------+
```

Each slot in the bucket holds both key and value.


---

### Hashing

When you do:

```go
m[key] = value
```

1. Go computes a **hash of the key**.

   * If key is a string, Go uses a **runtime hash function with a random seed** (to prevent DoS attacks).
   * If key is int or struct, Go uses appropriate hashing for its type.

2. Hash modulo **number of buckets** → gives the bucket index.

3. The key/value is stored in that bucket.

---

### Buckets & Overflow

Each bucket can store **8 key/value pairs** (fixed size).

* If more than 8 keys hash to the same bucket, **overflow buckets** are created.
* Overflow buckets are linked to the original bucket.

```
bucket0: key0,val0 | key1,val1 | ... | key7,val7
overflow → bucket0.1: key8,val8 | key9,val9 | ...
```

This handles **hash collisions** efficiently.

---

### Load Factor

**Load factor (α)** = number of elements ÷ number of buckets

* Go keeps load factor ≤ 6.5 (roughly)
* When load factor exceeds threshold, the map **resizes**

> Why?
> High load factor → more collisions → more overflow buckets → slower lookups

---

### Map Resizing (Grow)

When map grows:

1. Number of buckets doubles (`2^n` strategy)
2. All existing elements are **rehash into new buckets**

```
Old map: 8 buckets, 20 elements
New map: 16 buckets, 20 elements rehashed
```

* Go does **incremental rehashing**.
* Not all elements are moved at once; Go spreads rehashing over future operations.
* This prevents long pauses for very large maps.

---

### Why Incremental Resizing?

If Go moved **all elements at once**:

* For very large maps (millions of keys), inserting one element could block for milliseconds → noticeable latency
* Incremental approach spreads the cost → smoother runtime performance

---

# Memory Layout

### Each bucket:

* Fixed size: 8 key/value slots
* Each slot: key header + value header (24 bytes for slice or string on 64-bit)
* Overflow buckets: linked via pointer

```
+--------+--------+--------+--------+
| key0   | val0   | key1   | val1   | ... up to 8
+--------+--------+--------+--------+
```

* Map itself: pointer to bucket array + metadata

  * Number of buckets
  * Hash seed
  * Count of elements

---

# Performance Implications

1. **Lookup**: average O(1)
2. **Insertion**: average O(1)
3. **Worst-case**: O(n) if all keys collide (rare because of hash randomization)
4. **Iteration order**: unpredictable

   * Buckets + overflow + random hash seed → iteration order is intentionally randomized

---


# Concurrency Warning

Maps are **NOT thread-safe**.

This will panic:

```go
go func() { m["a"] = 1 }()
go func() { m["b"] = 2 }()
```

Error:

```
fatal error: concurrent map writes
```

---

### Solution 1: Use `sync.Mutex`

```go
import "sync"

var mu sync.Mutex
mu.Lock()
m["a"] = 1
mu.Unlock()
```

---

### Solution 2: Use `sync.Map`

```go
import "sync"

var sm sync.Map
sm.Store("a", 1)
v, ok := sm.Load("a")
```

Use `sync.Map` only when:

* Many goroutines
* Mostly reads
* Rare writes

---

# Passing Map to Function vs Reassigning

This works:

```go
func modify(m map[string]int) {
    m["x"] = 1
}
```

Maps (like slices) are **reference types** in Go, so modifying their contents inside a function affects the original.


This does NOT modify original:

```go
func replace(m map[string]int) {
    m = make(map[string]int)
}
```

another example

```go
func replaceMap(m map[string]int) {
    m = make(map[string]int) 
    m["y"] = 2
}

func main() {
    myMap := map[string]int{"x":1}
replaceMap(myMap)
fmt.Println(myMap) // still {"x":1}
}
```


However, **reassigning the map variable** inside a function only changes the local copy, leaving the original map unchanged. this is because When you pass a map to a function, the map variable itself is passed by value
Passed by value → copy of pointer; modifying data works, reassigning pointer does not.


Because:

* Map variable is passed by value
* Underlying data is reference
* Reassignment changes only local copy


You can also add new keys in the function — that modifies the underlying data, so it affects the caller:

```go
func addKey(m map[string]int) {
    m["new"] = 42
}

func main() {
    myMap := make(map[string]int)
    addKey(myMap)
    fmt.Println(myMap) // map[new:42]
}
```

Works perfectly because we’re modifying the data, not reassigning the map variable.



---

# Zero Value Behavior

If:

```go
var m map[string]int
```

Then:

* `m == nil`
* Reading is allowed
* Writing causes panic

```go
m["x"] = 10 // panic: assignment to entry in nil map
```

so Before writing, you need to initialize the map:

```go
m = make(map[string]int)  // allocate underlying hash table
m["x"] = 10               // now safe
```

---

# Capacity and Performance Tips

* Always use `make(map[K]V, size)` if size known.
* Avoid frequent resizing.
* Avoid using very large structs as keys.
* Prefer string or int keys for performance.

---

# Deep Copying a Map

Manual copy:

```go
newMap := make(map[string]int)
for k, v := range oldMap {
    newMap[k] = v
}
```

---

# Map vs Slice (When to Use)

Use map when:

* Fast lookup required
* Keys are meaningful
* Order doesn't matter

Use slice when:

* Order matters
* Indexed access required

---

Here’s a concise list of real-world backend use cases for Go maps:

1. **Caching / In-Memory Stores** – store frequently accessed data like user sessions or product info.
2. **Counting / Frequency Maps** – track occurrences of items, API hits, or sales.
3. **Lookup Tables / Configuration Maps** – map keys to handlers, error messages, or settings.
4. **Grouping / Indexing Data** – group orders by user, reviews by product, etc.
5. **Deduplication** – remove duplicate emails, IDs, or entries.
6. **Fast Authorization / Role Checks** – check user permissions or feature flags.
7. **Temporary State Storage** – keep transient state between requests in memory.
8. **Mapping Relationships** – e.g., user → list of friends, product → list of tags.
9. **Routing / Command Maps** – route HTTP endpoints or commands to functions.
10. **Feature Flags / Toggles** – map feature names to enabled/disabled status.

Maps are everywhere in backend because they give **fast, key-based access** and are very flexible.
