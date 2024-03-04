# Benchmark 🚀

Welcome to Benchmark, where I dive into the performance of different functions to help you make informed decisions!

## 📊 Benchmarked Functions

### Note:
- Unless explicitly mentioned, all benchmarks are done on a computer with **cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz**.

### Result Description
- **No of iterations:** Indicates the total number of times the loop was executed during the benchmark. 

-  **ns/op:** Indicates the average amount of time each iteration took to complete, expressed in nanoseconds per operation. This metric provides insights into the speed of the function execution.

- **Bytes/op:** Represents the average amount of memory allocated per operation. This metric quantifies the memory consumption of the process during execution.

- **Allocs/op:** Denotes the average frequency of memory allocations performed by the program while running the benchmark. It provides information about the allocation behavior of the function.

### 1. [Regular Expression vs. String ContainsAny](regExstring_test.go)

#### Results:

| Function             | No of iteration | ns/ops   | Bytes/op  | Allocs/op |
|----------------------|-----------------|----------|-----------|-----------|
| **RegEx**            | 744,342         | 1561     | 783       | 10
| **StringContainsAny**| 48,051,012      | 22.97    | 0         | 0
