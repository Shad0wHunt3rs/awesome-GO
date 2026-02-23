# Mutexes

Mutexes in Go are a **core synchronization primitive** used to protect shared data from **race conditions** when multiple goroutines access it concurrently.

## Why mutex is needed

Go supports **concurrency via goroutines**.

But concurrency introduces problems:

### Example (race condition)

```go
var counter int

func increment() {
    counter++
}
```

If multiple goroutines call `increment()`:

* Read counter
* Add 1
* Write back

These steps are **not atomic** → race occurs.

Result → incorrect value.

So we need:

> Only one goroutine should access critical section at a time

This is what a **mutex** does.

---

## What is a Mutex

Mutex = **Mutual Exclusion lock**

It ensures:

* Only one goroutine enters critical section
* Others wait

Go provides mutex in:

```go
sync.Mutex
```

From Go standard package:

```go
import "sync"
```

---

# Basic Mutex Usage

**Structure**

```go
type Counter struct {
    mu sync.Mutex
    value int
}
```

#### Critical section protection

```go
func (c *Counter) Inc() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}
```

#### Flow

```
Lock → critical section → Unlock
```

---

# Complete Example

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu sync.Mutex
    value int
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
    defer wg.Done()

    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

func main() {
    var c Counter
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go c.Inc(&wg)
    }

    wg.Wait()
    fmt.Println(c.value)
}
```

Output → always 1000

Without mutex → random.

---

# Internal working of Go Mutex

Go mutex is:

> Hybrid spin-lock + semaphore-based lock

### Fast path

If mutex free:

* Acquire using atomic CAS
* No kernel call

→ extremely fast

### Slow path

If locked:

* Spin briefly
* If still locked → goroutine sleeps
* Runtime scheduler wakes later

This design gives:

* Low latency
* High scalability

---

# Lock and Unlock rules

### Rule 1 — Unlock only if locked

```go
mu.Unlock() // panic if not locked
```

Runtime panic.

---

### Rule 2 — Same goroutine NOT required

Unlike some languages:

Go mutex is **not goroutine-owned**

Meaning:

* Goroutine A can Lock
* Goroutine B can Unlock

But this is bad design normally.

---

### Rule 3 — Must Unlock in all paths

Use defer:

```go
mu.Lock()
defer mu.Unlock()
```

But note:

* Defer adds small overhead
* OK in most cases

---

# TryLock (Go 1.18+)

Non-blocking lock attempt.

```go
if mu.TryLock() {
    // got lock
    mu.Unlock()
} else {
    // failed
}
```

Used when:

* Avoid blocking
* Opportunistic work

---

# RWMutex (Read-Write Mutex)

Very important advanced concept.

Problem:

* Many reads
* Few writes

Mutex blocks readers unnecessarily.

Solution:

```go
sync.RWMutex
```

Allows:

* Multiple readers
* Single writer

### Methods

```
RLock()
RUnlock()
Lock()
Unlock()
```

---

### Example

```go
var mu sync.RWMutex
var data int

func read() {
    mu.RLock()
    fmt.Println(data)
    mu.RUnlock()
}

func write() {
    mu.Lock()
    data++
    mu.Unlock()
}
```

---

# RWMutex rules

### Rule 1

If writer waiting → new readers blocked

Prevents writer starvation.

---

### Rule 2

Do NOT upgrade lock

This is wrong:

```go
mu.RLock()
mu.Lock() // DEADLOCK
```

Must release first.

---

# Mutex vs Channels

Go philosophy:

> Do not communicate by sharing memory; share memory by communicating

But reality:

### Use channels when

* Ownership transfer
* Pipelines
* Message passing

### Use mutex when

* Shared state
* Cache
* Counters
* In-memory structures

Senior Go engineers use both.

---

# Common Mutex patterns

## Pattern 1 — Embedded mutex

```go
type Cache struct {
    sync.Mutex
    m map[string]int
}
```

Then:

```go
c.Lock()
```

But explicit field preferred for readability.

---

## Pattern 2 — Guarded struct

Mutex protects entire struct.

Most common.

---

## Pattern 3 — Fine-grained locking

Multiple mutexes:

* Improves concurrency
* Harder design

Used in high-performance systems.

---

# Deadlocks

### Self-deadlock

```go
mu.Lock()
mu.Lock() // deadlock
```

Mutex not reentrant.

---

### Cross deadlock

```go
A locks mu1
B locks mu2

A waits mu2
B waits mu1
```

Classic.

Solution:

* Lock ordering

---

# Copying mutex

This is wrong:

```go
type X struct {
    mu sync.Mutex
}

x2 := x1 // copies mutex
```

Mutex must NOT be copied after use.

Because:

* Internal state duplicated
* Undefined behavior

---

# Mutex performance tips

### Tip 1

Keep critical section small.

### Tip 2

Avoid blocking operations inside lock.

Bad:

```go
mu.Lock()
time.Sleep(1*time.Second)
mu.Unlock()
```

---

### Tip 3

Avoid IO inside lock.

---

# Advanced runtime insight (high-level)

Go runtime uses mutex internally for:

* Scheduler queues
* Heap allocator
* Map growth
* Channel synchronization

Mutex is fundamental runtime primitive.

---

# When NOT to use mutex

Avoid if:

* Pure message passing
* Actor model
* Immutable data
* Single goroutine ownership

---
