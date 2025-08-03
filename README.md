# go-telemetry

A lightweight Go wrapper for emitting application metrics and logs. This package provides abstractions over Prometheus and console output for metrics, and structured logging for log messages. It simplifies observability setup by allowing flexible, environment-based switching of providers.

---

## ✨ Features

- ✅ Unified interface for metrics emission and logging
- 📊 Prometheus histogram and counter support
- 🖨️ Console metrics emitter for local development
- ⏱️ Built-in latency tracking for methods
- 📄 Structured logging with service/method context
- 🔁 Configurable via environment variables

---

## 📦 Installation

```bash
go get github.com/hladush/go-telemetry
```


### 🧩 Environment Variables
| Variable             | Description                                                                   | Default   |
| -------------------- | ----------------------------------------------------------------------------- | --------- |
| `METRICS_EMITTER`    | Backend for metrics:`noop`, `prometheus` or `console`                         | `console` |
| `PROMETHEUS_ADDRESS` | Address to expose Prometheus metrics endpoint (used by `http.ListenAndServe`) | `:8081`   |


### 🚀 Usage Example

Here's a quick example of how to use `go-telemetry` in your Go application:

```go
import (
    "errors"
    "net/http"
    "time"

    "github.com/hladush/go-telemetry"
)

var (
    serviceName   = "MainHandler"
    methodMetrics = telemetry.NewMethod("MyMethod", serviceName)
)

func myMethod(w http.ResponseWriter, r *http.Request) {
    // Track method latency
    defer methodMetrics.RecordLatency(time.Now(), "in", "getDevices")

    // Increment a counter for method calls
    methodMetrics.IncCounter()

    // Log and count an error occurrence
    methodMetrics.LogAndCountError(errors.New("failed to get devices"), "marshal")

    // Count specific error types
    methodMetrics.CountError("success", "marshal")
    methodMetrics.CountError("another_error")

    // Log and count a successful operation
    methodMetrics.LogAndCountSuccess("success", "in", "getDevices")

    // Increment a custom counter with labels
    methodMetrics.IncCounter("inc_counter", "in", "getDevices")
}
```
