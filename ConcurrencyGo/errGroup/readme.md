Here’s a clean and professional `README.md` for your project.

You can copy-paste this directly into your repository.

---

# Concurrent CSV Reader in Go

A demonstration of concurrent file processing patterns in Go using:

* `sync.WaitGroup`
* `errgroup.Group`
* `errgroup.WithContext`

This project shows how concurrency evolves from basic synchronization to production-grade cancellation-aware execution.

---

## 📌 Overview

The program reads multiple CSV files concurrently:

```
file1.csv
file2.csv
file3.csv
```

Each file is processed in its own goroutine and the records are printed to stdout.

The repository demonstrates three approaches:

1. **WaitGroup** – Basic synchronization
2. **errgroup.Group** – Error aggregation
3. **errgroup.WithContext** – Error handling + cancellation (recommended)

---

## 🧠 Concepts Covered

* Goroutines
* Channels
* Fan-out concurrency pattern
* `sync.WaitGroup`
* `golang.org/x/sync/errgroup`
* Context cancellation
* Cooperative shutdown
* Fail-fast concurrency design

---

## 🏗 Project Structure

```
.
├── main.go
├── file1.csv
├── file2.csv
├── file3.csv
└── README.md
```

---

## 🚀 How It Works

### 1️⃣ WaitGroup Version

* Spawns one goroutine per file
* Uses `sync.WaitGroup` to wait for completion
* Errors are printed but do not stop other goroutines

**Characteristics:**

* Manual synchronization
* No cancellation
* All workers continue even if one fails

---

### 2️⃣ errgroup.Group Version

* Uses `errgroup.Group`
* Collects errors from goroutines
* Returns the first error encountered

**Characteristics:**

* Cleaner error handling
* Still no cancellation
* Other goroutines continue running

---

### 3️⃣ errgroup.WithContext Version (Production Style)

* Uses `errgroup.WithContext`
* Automatically cancels all goroutines on first error
* Supports timeouts
* Implements cooperative cancellation via `ctx.Done()`

```go
select {
case <-ctx.Done():
    return ctx.Err()
case line := <-ch:
}
```

**Characteristics:**

* Fail-fast behavior
* Graceful shutdown
* Timeout support
* Production-ready concurrency model

---

## 🔥 Why `errgroup.WithContext` Is Preferred

In real-world systems:

* If one worker fails, others should stop.
* Long-running tasks should respect cancellation.
* Systems should support timeouts.

`errgroup.WithContext` provides all of this cleanly.

---

## ▶️ Running the Project

Make sure you have Go installed (1.20+ recommended).

Initialize module if needed:

```bash
go mod init concurrent-csv
go get golang.org/x/sync/errgroup
```

Run:

```bash
go run main.go
```

---

## 🧪 Optional: Test Cancellation

Uncomment this in `main()`:

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
defer cancel()
```

This forces early cancellation and demonstrates cooperative shutdown.

---

## 📊 Concurrency Evolution

| Feature                | WaitGroup | errgroup | errgroup + Context |
| ---------------------- | --------- | -------- | ------------------ |
| Concurrent execution   | ✅         | ✅        | ✅                  |
| Error propagation      | ❌         | ✅        | ✅                  |
| Automatic cancellation | ❌         | ❌        | ✅                  |
| Timeout support        | ❌         | ❌        | ✅                  |
| Production ready       | ⚠️        | ⚠️       | ✅                  |

---

## 🎯 Key Takeaways

* `WaitGroup` is good for simple synchronization.
* `errgroup` improves structured error handling.
* `errgroup.WithContext` is the correct approach for production systems.
* Always design concurrent systems with cancellation in mind.

---
