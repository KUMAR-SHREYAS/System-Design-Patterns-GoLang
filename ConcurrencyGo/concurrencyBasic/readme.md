

# Go Concurrency Basics – Buffered Channels, Goroutines & Graceful Shutdown

This project demonstrates core Go concurrency concepts:

- Goroutines
- Buffered channels
- Channel closing
- Detecting closed channels
- Producer–Consumer pattern
- Graceful shutdown using signal channels
- Concurrency vs Parallelism



## 📌 Program Overview

We use:

- A **producer goroutine** that sends integers
- A **consumer goroutine** that receives them
- A buffered channel to control flow
- A signaling channel to gracefully exit `main()`

---

## 🧠 Key Code Structure

### Channel Creation

```go
ch := make(chan int, 2)
exit := make(chan struct{})
````

* `ch` → Buffered channel with capacity 2
* `exit` → Used as a completion signal




---
### Producer Goroutine

```go
go func() {
    for i := 0; i < 5; i++ {
        fmt.Println(time.Now(), i, "sending")
        ch <- i
        fmt.Println(time.Now(), i, "sent")
        time.Sleep(1 * time.Second)
    }

    fmt.Println(time.Now(), "all completed, leaving")
    close(ch)
}()
```
### What It Does

* Sends numbers 0 → 4
* Sleeps 1 second between sends
* Closes channel after finishing

---

## 🔁 Consumer Goroutine

```go
go func() {
    for {
        select {
        case v, open := <-ch:
            if !open {
                close(exit)
                return
            }
            fmt.Println(time.Now(), "received", v)
        }
    }
}()
```

### What It Does

* Continuously reads from channel
* Detects when channel is closed
* Signals `exit` channel when done

---

## ⏳ Main Goroutine

```go
fmt.Println(time.Now(), "waiting for everything to complete")
<-exit
fmt.Println(time.Now(), "exiting")
```

Main blocks until:

* Consumer detects closed channel
* Consumer closes `exit`
* Main resumes and exits cleanly

---

# 🔥 Buffered Channel Behavior

When you create:

```go
make(chan int, 2)
```

It means:

* Capacity = 2
* Can hold 2 values before blocking sender
* FIFO queue behavior

### Example

| Operation | Buffer State | Size     |
| --------- | ------------ | -------- |
| send 0    | [0]          | 1        |
| send 1    | [0,1]        | 2 (full) |
| receive   | [1]          | 1        |
| send 2    | [1,2]        | 2        |

### Important Rules

* If buffer is **full → sender blocks**
* If buffer is **empty → receiver blocks**
* When receiver reads → buffer size decreases by 1

---

# 🔍 Understanding `v, open := <-ch`

When reading from channel:

```go
v, open := <-ch
```

* `v` → value received
* `open` → false if channel is closed

If `open == false`:

* No more values will arrive
* Safe to terminate consumer

---

# 🧵 Concurrency vs Parallelism

## Concurrency

Multiple goroutines are in progress during overlapping time.

## Parallelism

Multiple goroutines execute at the exact same time on different CPU cores.

Your program is:

✔ Concurrent
✔ Possibly parallel (depends on CPU cores & GOMAXPROCS)

Even on one CPU core, Go scheduler switches between goroutines.

---

# ⚙️ How Go Scheduler Works (High Level)

* Goroutines are lightweight threads
* Go uses an M:N scheduler
* Many goroutines multiplexed over fewer OS threads
* Work-stealing scheduling model
* Controlled by `GOMAXPROCS`

Check CPU parallelism:

```go
runtime.GOMAXPROCS(0)
```

---

# 🧹 Why This Program Shuts Down Cleanly

1. Producer closes `ch`
2. Consumer detects closed channel
3. Consumer closes `exit`
4. Main waits on `exit`
5. Program exits gracefully

No goroutine leaks
No deadlocks
Proper lifecycle management

---

# 🧠 Cleaner Alternative (Without Select)

Since only one channel is used, consumer could be simplified:

```go
go func() {
    for v := range ch {
        fmt.Println("received", v)
    }
    close(exit)
}()
```

`range` automatically exits when channel is closed.

---

# 🎯 Concepts Demonstrated

* Goroutines
* Buffered channels
* Blocking behavior
* Channel closing rules
* Signaling with `struct{}`
* Producer–Consumer pattern
* Graceful shutdown pattern
* Concurrency fundamentals

---

# 🚀 Interview Takeaways

If asked in interview:

* Only sender should close channel
* Never close channel from receiver
* Use `struct{}` for signaling
* Use `range` when single channel
* Use `select` when multiple channels
* Buffered channel reduces blocking but doesn’t eliminate it

---

# 🏁 Expected Output Pattern

Example:

```
waiting for everything to complete
0 sending
0 sent
received 0
1 sending
1 sent
received 1
...
all completed, leaving
exiting
```

Exact timestamps may vary.

---

# 📌 Final Summary

This example demonstrates safe and structured concurrency in Go using:

* Buffered channels
* Proper closing strategy
* Goroutine coordination
* Graceful termination

It is a foundational concurrency pattern in Go systems.

---

