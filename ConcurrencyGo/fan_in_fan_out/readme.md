Here’s a clean, professional `README.md` for your **Fan-In / Fan-Out Go implementation**.

You can directly copy this into your repository.

---

# Fan-In & Fan-Out Pattern in Go

This project demonstrates the **Fan-Out** and **Fan-In** concurrency patterns in Go using goroutines and channels.

It showcases how to:

* Distribute work across multiple workers (Fan-Out)
* Merge results from multiple channels into one (Fan-In)
* Coordinate concurrent processing safely and efficiently

---

## 📌 Overview

The program reads multiple CSV files concurrently and merges their outputs into a single stream.

It demonstrates real-world concurrency design patterns commonly used in:

* Distributed systems
* Streaming pipelines
* Worker pools
* Microservices
* High-performance backends

---

# 🧠 Concepts Covered

* Goroutines
* Channels
* Buffered vs unbuffered channels
* Fan-Out pattern
* Fan-In pattern
* Synchronization with `sync.WaitGroup`
* Concurrent file processing

---

# 🔥 What is Fan-Out?

Fan-Out means:

> One input → Multiple workers processing in parallel

In your implementation:

```go
for _, file := range files {
    go read(file)
}
```

Each file is processed in its own goroutine.

### Why use Fan-Out?

* Parallel execution
* Better CPU utilization
* Faster I/O processing
* Scalable workload distribution

---

# 🔥 What is Fan-In?

Fan-In means:

> Multiple inputs → One combined output channel

In your implementation:

```go
func fanIn(chans ...<-chan []string) <-chan []string
```

This function:

* Accepts multiple channels
* Launches goroutines to read from each
* Forwards all data into a single output channel

---

# 🏗 Architecture Flow

```
             file1.csv   ┐
             file2.csv   ├──> Fan-Out (multiple goroutines)
             file3.csv   ┘
                    ↓
             Individual channels
                    ↓
               Fan-In
                    ↓
           Single merged channel
                    ↓
                main()
```

---

# 🛠 How It Works

## 1️⃣ Fan-Out Phase

Each file:

* Opens independently
* Reads records
* Sends records to its own channel

```go
func read(file string) <-chan []string
```

Each invocation runs inside a goroutine.

---

## 2️⃣ Fan-In Phase

The `fanIn` function:

* Takes multiple channels
* Uses `sync.WaitGroup`
* Forwards all messages to one output channel
* Closes output only after all inputs finish

```go
func fanIn(chans ...<-chan []string) <-chan []string
```

This ensures:

* No data loss
* Proper synchronization
* Clean shutdown

---

# ▶️ Running the Project

Ensure Go 1.20+ is installed.

```bash
go mod init fan-in-fan-out
go run main.go
```

Make sure the CSV files exist:

```
file1.csv
file2.csv
file3.csv
```

---

# ⚙️ Key Implementation Detail

### Why use `file := file` inside loop?

```go
for _, file := range files {
    file := file
    go func() { ... }()
}
```

This prevents closure capture bugs where all goroutines reference the same loop variable.

---

# 📊 Benefits of Fan-In / Fan-Out

| Feature            | Benefit                     |
| ------------------ | --------------------------- |
| Parallelism        | Faster processing           |
| Scalability        | Easy to add workers         |
| Isolation          | Failures don’t block others |
| Clean coordination | Safe channel closing        |

---

# 🧪 Real-World Use Cases

* Log aggregation systems
* Distributed file processing
* API request fan-out calls
* Search engine indexing
* Event streaming pipelines
* Worker pool architectures

---

# 🎯 Why This Pattern Matters

Fan-In / Fan-Out is a foundational concurrency model in Go.

It enables:

* Clean pipeline architecture
* High throughput systems
* Structured concurrent design
* Backpressure handling via channels

---
