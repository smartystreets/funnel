# funnel

A generic fan-out/fan-in library in Go.

Allows for easy implementation of the fan-out/fan-in pattern described on the [Go blog](https://go.dev/blog/pipelines).

How to use:

1. Load up (and CLOSE!) an input channel (optionally in a separate goroutine).
2. Provide the desired number of workers, the input channel, and a function that will process each item from the input channel to the FanOut function.
3. Read processed items from the channel returned by the FanOut function until it is consumed.

See the test file for an example.

Notes:

- Responsibility to close the input channel rests with the caller.
- Responsibility to close the output channel rests with this library.
- Passing `0` as the number of workers will result in a panic.
- While each item provided to the input channel will be processed, the order of those items on the output channel is non-deterministic.
- Calling FanOut will result in `2N+1` goroutines, where N is the number of workers specified.
- Each goroutine created will have exited by the time the output channel has been consumed.
