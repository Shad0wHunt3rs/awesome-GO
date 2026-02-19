# Concurrency vs Parallelism

## Concurrency

**Concurrency = dealing with multiple tasks at once (structure)**

* Multiple tasks make progress over time
* Tasks may be interleaved on a single CPU
* Focus is **coordination**

**Example**

One CPU:

```
Task A → Task B → Task A → Task C → Task B ...
```

All tasks progress → concurrent.

---

## Parallelism

**Parallelism = executing multiple tasks at the same time (hardware)**

* Requires multiple CPU cores
* Focus is **simultaneous execution**

**Example**

4 cores:

```
Core1: Task A
Core2: Task B
Core3: Task C
Core4: Task D
```

---

## Key Insight

> Concurrency is about **design**
> Parallelism is about **execution**

You can have:

* concurrency without parallelism (single core)
* parallelism requires concurrency

---

# Why Go is Designed for Concurrency

Go was created at Google where systems needed:

* network servers
* distributed systems
* microservices
* I/O heavy workloads

Traditional threads were:

* heavy
* complex
* expensive to create
* hard synchronization

Go introduced:

* goroutines (cheap threads)
* channels (communication)
* CSP model (communicate instead of share memory)

---

# Go Concurrency Model (CSP)

Go follows **Communicating Sequential Processes (CSP)**.

The principle:

> Do not communicate by sharing memory; share memory by communicating.

Meaning:

* Instead of locks
* Use channels

---

# Goroutines (Lightweight Threads)

A goroutine is a **function executing concurrently with others**.

**Syntax**

```go
go function()
```

Example:

```go
go fmt.Println("hello")
```

---

## Goroutine Characteristics

### 1. Very cheap

* Thread: ~1–2 MB stack
* Goroutine: ~2 KB stack (grows dynamically)

You can create millions.

---

### 2. M:N Scheduling

Go runtime maps:

```
M goroutines → N OS threads
```

This is called **green threading model**.

---

### 3. Independent stack

Each goroutine has:

* its own stack
* but shared heap

---

### 4. Non-blocking creation

Creating goroutine does not block caller.

---

## ⚠️ Important beginner issue

```go
func main() {
    go fmt.Println("hello")
}
```

Output may not appear → program exits.

Fix:

```go
time.Sleep(time.Second)
```

* main() starts
* go fmt.Println("hello") creates a goroutine
* Goroutine is scheduled (not executed immediately)
* main() finishes instantly
* Program exits
* Runtime kills all goroutines

👉 So the goroutine never gets CPU time.

---

# 5. Go Scheduler (Deep Understanding)

The Go scheduler is the heart of concurrency.

It uses **G-M-P model**:

| Component | Meaning                          |
| --------- | -------------------------------- |
| G         | Goroutine                        |
| M         | Machine (OS thread)              |
| P         | Processor (logical CPU resource) |

---

## Flow

```
G → P → M → CPU
```

* G needs P to run
* P attached to M
* M runs on CPU

---

## Why P exists?

P holds:

* run queue
* scheduler state
* memory cache

This improves performance.

---

## Work stealing

If one P becomes idle:

* it steals goroutines from another P

This balances load.

---

# GOMAXPROCS (Parallelism control)

`GOMAXPROCS` controls number of CPUs used simultaneously.

Default:

```
= number of cores
```

Example:

```go
runtime.GOMAXPROCS(1)
```

Now:

* concurrency exists
* parallelism disabled

---

# Channels (Communication Primitive)

Channel = typed pipe between goroutines.

**Creation**

```go
ch := make(chan int)
```

---

## Send / Receive

```go
ch <- 10   // send
x := <-ch  // receive
```

---

## Blocking behavior

| Operation | Behavior                    |
| --------- | --------------------------- |
| Send      | blocks until receiver ready |
| Receive   | blocks until sender ready   |

This is **synchronization by design**.

---

# Buffered Channels

```go
ch := make(chan int, 3)
```

Now:

* send blocks only when buffer full
* receive blocks when empty
* The 3 is the channel capacity

```go
ch := make(chan int, 3)

ch <- 1   // ok (slot 1)
ch <- 2   // ok (slot 2)
ch <- 3   // ok (slot 3)
ch <- 4   // BLOCKS (buffer full)
```

---

# Channel Closing

```go
close(ch)
```

Receiver can detect:

```go
v, ok := <-ch
```

* ok = false → closed

---

# Select Statement (Multiplexing)

Used to wait on multiple channel operations.

```go
select {
case x := <-ch1:
case ch2 <- 10:
default:
}
```

Equivalent of:

* epoll
* poll
* multiplexed wait

---

# Synchronization Primitives

## WaitGroup

Wait for goroutines to finish.

```go
var wg sync.WaitGroup
wg.Add(1)

go func(){
    defer wg.Done()
}()

wg.Wait()
```

---

## Mutex

Protect shared memory.

```go
var mu sync.Mutex

mu.Lock()
x++
mu.Unlock()
```

---

## Rule

* channels → coordination
* mutex → shared state protection

---

# Race Conditions

Occurs when:

* multiple goroutines access shared data
* at least one write
* no synchronization

Example:

```go
x++
```

Fix:

* mutex
* channel
* atomic

---

## Detection

```bash
go run -race main.go
```

---

# Atomic Operations

Low-level lock-free operations.

```go
atomic.AddInt64(&x, 1)
```

Faster than mutex for simple counters.

---

# Context Package (Cancellation)

Used for:

* timeouts
* cancellation
* request lifecycle

Example:

```go
ctx, cancel := context.WithCancel(context.Background())
```

---

# Patterns in Go Concurrency

## Worker pool

Multiple workers process jobs.

## Pipeline

Stage1 → Stage2 → Stage3

## Fan-out / Fan-in

* distribute work
* collect results

## Producer-consumer

Classic channel use.

---

# Concurrency vs Parallelism in Go (Final mental model)

## Concurrency in Go

* goroutines
* channels
* select
* WaitGroup
* mutex
* context

## Parallelism in Go

* scheduler
* multiple P
* multiple cores
* GOMAXPROCS

---

# Interview-Level Summary

**Concurrency**

* structure
* goroutines
* channels
* coordination
* scheduler interleaving

**Parallelism**

* hardware
* simultaneous execution
* multiple cores
* GOMAXPROCS

---

