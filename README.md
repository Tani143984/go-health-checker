# 🚀 Concurrent URL Health Checker (Go)

A high-performance command-line tool written in **Go** designed to monitor multi-endpoint availability and measure network latency concurrently. 

This project demonstrates core backend systems patterns and idiomatic Go concurrency primitives.

---

## 🛠️ Key Features
* **Parallel Execution:** Leverages Go **Goroutines** to perform non-blocking concurrent HTTP checks across multiple endpoints.
* **Concurrency Synchronization:** Utilizes **`sync.WaitGroup`** to coordinate background workers and ensure safe execution completion before process exit.
* **Safe Data Flow:** Funnels asynchronous worker outcomes safely via **Buffered Channels**.
* **Resilience & Timeouts:** Implements custom `http.Client` timeouts and explicit error-state validation to prevent execution hangs on unreachable hosts.
* **Structured JSON Reporting:** Serializes execution results into clean, indented JSON diagnostics.

---

## ⚙️ Tech Stack
* **Language:** Go (Golang) 1.22+
* **Standard Libraries:** `net/http`, `encoding/json`, `sync`, `time`, `os`

---

## 🏃‍♂️ Quick Start

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/Tani143984/go-health-checker.git](https://github.com/Tani143984/go-health-checker.git)
   cd go-health-checker
