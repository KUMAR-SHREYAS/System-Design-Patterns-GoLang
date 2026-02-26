
# 🚀 Concurrency Pipeline Pattern in Go

## 📌 Overview

This repository demonstrates the **Pipeline Concurrency Pattern** in Go — a structured approach for processing data in stages using goroutines and channels.

The pipeline pattern is widely used in high-performance systems where data flows through multiple transformation or filtering stages.

---

## 🧠 What is the Pipeline Pattern?

A pipeline consists of:

* **Independent stages**
* Each stage runs in its own **goroutine**
* Stages communicate via **channels**
* Output of one stage becomes input of the next

### Conceptual Flow

```
Source → Stage 1 → Stage 2 → Stage 3 → Output
```

Each stage:

* Receives input from a channel
* Processes data
* Sends results to another channel
* Closes the channel when done

---

## 🏗 Architecture

```
+------------+      +------------+      +------------+      +---------+
|  Producer  | ---> |  Processor | ---> |  Filter    | ---> | Consumer|
+------------+      +------------+      +------------+      +---------+
       |                   |                   |                 |
   Goroutine           Goroutine           Goroutine         Main Thread
```

---

## ⚙ Key Concepts Used

* **Goroutines** – Lightweight concurrent functions
* **Channels** – Communication between stages
* **Range over channel** – Streaming data processing
* **Channel closing** – Proper pipeline termination
* **Read-only channels (`<-chan`)** – Safer APIs

---

## 🎯 Why Use Pipeline Concurrency?

### ✅ Streaming Processing

Data is processed as it arrives — no need to load everything into memory.

### ✅ Separation of Concerns

Each stage has a single responsibility.

### ✅ Scalability

Stages can be parallelized (fan-out/fan-in).

### ✅ Clean Error Handling

Can be combined with `errgroup` or context cancellation.

---

## 🔄 Typical Use Cases

* ETL (Extract → Transform → Load) systems
* Log processing
* File processing
* Streaming data systems
* Kafka consumers
* ML preprocessing pipelines
* Network packet handling
* Batch job systems

---

## 📦 Example Pipeline Structure

```go
func stage1(in <-chan Data) <-chan Data
func stage2(in <-chan Data) <-chan Data
func stage3(in <-chan Data) <-chan Result
```

Usage:

```go
out := stage3(stage2(stage1(input)))
```

---

## ⚠ Best Practices

✔ Always close output channels
✔ Never close a channel you did not create
✔ Use read-only channel types for safety
✔ Handle errors properly
✔ Use `context.Context` for cancellation
✔ Consider worker pools for CPU-bound stages

---

## 🧩 Advanced Extensions

* **Fan-Out / Worker Pools**
* **Fan-In Aggregation**
* **Context Cancellation**
* **Backpressure Handling**
* **Bounded Channels**
* **errgroup Integration**
* **Rate Limiting**

---

## 📊 Comparison With Sequential Processing

| Sequential          | Pipeline                     |
| ------------------- | ---------------------------- |
| Single thread       | Multiple concurrent stages   |
| Blocks on each step | Stages operate independently |
| Higher latency      | Lower latency                |
| Not scalable        | Easily parallelizable        |

---

## 🧵 When NOT to Use Pipeline

* Extremely small workloads
* Simple linear logic
* When ordering must be strictly preserved without coordination
* When concurrency overhead outweighs benefit

---

## 📚 Learning Goals

This project helps understand:

* Go concurrency model
* Channel synchronization
* Streaming architecture
* Building scalable backend systems

---

## 🔥 Inspiration

The pipeline pattern is inspired by:

* Unix pipelines (`|`)
* Assembly line processing
* Reactive streams
* Distributed data processing systems

---
