# funnel

> fun + go channel = funnel

Just a series of exercises in understanding go concurrency.

## Concepts

1. Goroutines are concurrently executing functions (within the same address space)
2. The main function of a Go program runs in a goroutine (running programs always have at least one goroutine)
3. When the 'main' goroutine exits, the entire process exits (along with any unfinished goroutines)
4. The Go scheduler decides which goroutine(s) runs at any given moment, but they are limited in how they transition from one goroutine to another
5. When a goroutine encounters a blocking I/O operation the scheduler is given the opportunity to transition to a different goroutine
6. When no goroutine can proceed a deadlock occurs (which results in a panic)
7. Because goroutines run in the same address space, they can refer to variables with common scope
8. When multiple goroutines attempt to access the same variable and at least one access is a 'write' operations, there is a potential for a data race, which could result in crashes or memory corruption
9. One approach to safe concurrency is to use a 'mutex' (mutual exclusion lock)
10. Another approach to safe concurrency is Go channels, which are FIFO queue structures that are actually safe to send to/receive from across multiple goroutines
11. Go Slogan: "Do not communicate by sharing memory; instead, share memory by communicating."
12. Attempting to receive on an empty channel blocks the current goroutine until a value is sent or the channel is closed (which results in the 'zero value' being received).
13. Attempting to send on an 'unbuffered' channel blocks until the value is received or the channel is closed (which then results in a panic)
14. Concurrency != Parallelism
15. A Go program that is limited to running a single goroutine at a time is said to be concurrent
16. A Go program that can run multiple goroutines at the same time (such as by multiple CPUs) is said to be parallel

## Resources

- The 'go' Statement: https://go.dev/ref/spec#Go_statements
- Channel Types: https://go.dev/ref/spec#Channel_types
- Race Detector: https://go.dev/doc/articles/race_detector
- Concurrency: https://go.dev/doc/effective_go#concurrency
- Concurrency Patterns - Pipelines: https://go.dev/blog/pipelines
- Concurrency Patterns - Context: https://go.dev/blog/context
- Concurrency Patterns - Timing out, moving on: https://go.dev/blog/concurrency-timeouts (probably obsolete since advent of 'context')
- Advanced Concurrency Patterns: https://go.dev/blog/io2013-talk-concurrency
- Concurrency is not Parallelism: https://go.dev/blog/waza-talk
- Channel Axioms: https://dave.cheney.net/2014/03/19/channel-axioms
- Visualizing Go Concurrency: https://divan.dev/posts/go_concurrency_visualize/
