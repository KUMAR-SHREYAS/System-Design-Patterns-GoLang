

# PubSub Concurrency Pattern in Go

A minimal, thread-safe implementation of the **Publish–Subscribe (PubSub)** concurrency pattern using Go generics, channels, and `sync.RWMutex`.

This project demonstrates how to:

* Create a generic PubSub system
* Allow multiple subscribers
* Publish messages concurrently
* Gracefully close all subscribers
* Coordinate goroutines using `sync.WaitGroup`

---

## 📌 What is the PubSub Pattern?

The **Publish–Subscribe pattern** is a messaging pattern where:

* **Publishers** send messages
* **Subscribers** receive messages
* Publishers and subscribers are **decoupled**
* A central broker (PubSub) manages distribution

In Go, this is naturally implemented using **channels and goroutines**.

---

## 🏗 Architecture Overview

```
Publisher
    |
    v
  PubSub  ----> Subscriber 1 (goroutine)
    |          
    └---------> Subscriber 2 (goroutine)
```

Core components:

* `Publish(value)` → Broadcasts value to all subscribers
* `Subscribe()` → Returns a read-only channel
* `Close()` → Closes all subscriber channels safely

---

## 🧠 Implementation Breakdown

### PubSub Structure

```go
type PubSub[T any] struct {
	subs   []chan T
	mu     sync.RWMutex
	closed bool
}
```

### Fields Explained

| Field    | Purpose                            |
| -------- | ---------------------------------- |
| `subs`   | Stores all subscriber channels     |
| `mu`     | Ensures thread-safe access         |
| `closed` | Prevents publishing after shutdown |

---

## 🔐 Concurrency Safety

This implementation uses:

* `sync.RWMutex`

  * `RLock()` for publishing (multiple readers allowed)
  * `Lock()` for subscribing and closing (exclusive access)
* `sync.WaitGroup`

  * Ensains goroutines exit before program termination

This prevents:

* Race conditions
* Writing to closed channels
* Data corruption

---

## 🚀 Example Usage

### Creating PubSub

```go
ps := NewPubSub[string]()
```

### Subscribing

```go
s1 := ps.Subscribe()
```

Each subscriber receives its own channel.

### Publishing

```go
ps.Publish("one")
ps.Publish("two")
ps.Publish("three")
```

All subscribers receive all published values.

### Closing

```go
ps.Close()
```

This:

* Closes all subscriber channels
* Signals goroutines to exit gracefully

---

## 🖥 Sample Output

```
sub 1, value one
sub 2, value one
sub 1, value two
sub 2, value two
sub 1, value three
sub 2, value three
sub 1, exiting
sub 2, exiting
completed
```

---

## ⚙️ Key Concurrency Concepts Demonstrated

### 1️⃣ Generic Type Support

```go
type PubSub[T any]
```

Works with any type: `string`, `int`, structs, etc.

---

### 2️⃣ Fan-Out Pattern

`Publish()` distributes one message to multiple subscribers.

```go
for _, ch := range s.subs {
	ch <- value
}
```

---

### 3️⃣ Graceful Shutdown

Channels are closed inside `Close()`:

```go
for _, ch := range s.subs{
	close(ch)
}
```

Subscribers detect closure using:

```go
val, ok := <-ch
if !ok {
	return
}
```

---

## ⚠️ Important Considerations

### Blocking Behavior

This implementation uses **unbuffered channels**:

```go
r := make(chan T)
```

That means:

* `Publish()` will block until every subscriber receives the value.
* A slow subscriber can block the entire system.

### Production Improvements

For real systems, consider:

* Buffered channels
* Non-blocking publish (using `select`)
* Subscriber removal support
* Context cancellation
* Backpressure handling

---

## 📦 When to Use This Pattern

Ideal for:

* Event broadcasting
* Notification systems
* In-memory messaging
* Microservice internal signaling
* Real-time updates

Not ideal for:

* Persistent messaging
* Distributed systems (use Kafka, NATS, etc.)

---

## 🧪 How to Run

```bash
go run main.go
```

---

## 🏁 Learning Objectives

After studying this implementation, you should understand:

* How Go channels enable message broadcasting
* How `sync.RWMutex` protects shared state
* How goroutines coordinate using `WaitGroup`
* The trade-offs of blocking vs buffered PubSub designs

---

## 📚 Concepts Covered

* Goroutines
* Channels
* Generics (Go 1.18+)
* RWMutex
* WaitGroup
* Graceful shutdown
* Fan-out concurrency pattern

---

