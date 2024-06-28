# When to use concurrency

https://www.youtube.com/watch?v=YHRO5WQGh0k

- Apply competitiveness only when it makes your program better

- Apply competitiveness,
  when you need to combine the results of multiple operations that may be
  that can be performed independently of each other.

# All great to know things

- You can run any function as a goroutine

- Unbuffered channels have
  a very useful property: the sending side is blocked until the receiving coprogram reads data from the channel,
  until the receiving coprogram reads data from the channel, and vice versa, the receiving side is blocked until the sending coprogram
  does not write the value into the channel

- When transmitting an unbuffered channel, the len and cap functions return 0.
  This is quite logical, because by definition an unbuffered channel has no buffer in which to place values.
  has no buffer in which to place the values of the

- In most cases, unbuffered channels should be used

- When reading from a channel that may be closed, always
  use the comma-ok idiom to ensure that the channel is still open.
  open

- Channel should be closed only if
  a certain goroutine expects to close the channel (like, for example, a goroutine,
  reading from a channel using a for-range loop). Since channels are just another kind of variable, the Go language's runtime environment can detect already unused variables.
  Go language runtime can detect channels that are already unused and remove them by garbage collection.
  them by garbage collection

  - In most cases, you should not use the default branch inside a loop
    for-select. With this approach, the default branch will be triggered at each iteration of the loop.
    iteration of the loop that has not found case branches with executable operations
    read or write operations. As a result, the for loop will run continuously, increasing the CPU time consumption.

- Using range with channels in Go is the preferred way to iterate over channel elements. Here are a few reasons why this is so:

Simple and safe: range automatically handles reading from a channel before closing it. This prevents the program from locking on reading from the channel when there is no more data.

Performance: The use of range is optimized by the Go compiler and ensures efficient channel handling. The compiler takes into account the internal logic of channel operation to minimize overhead.

# When to use buffered and unbuffered channels

- By default, Go uses unbuffered channels,
  the principle of operation of which is quite simple: one goroutine makes a record
  and waits for the other one to "pick up" the result of its work, like a baton.
  like a baton. It is much more difficult to deal with buffered channels.
  more difficult. You have to choose the buffer size, because a buffered channel cannot have an unlimited buffer size.
  cannot have an unlimited buffer size. To use a buffered channel correctly, you must also provide for handling the case when the buffer is filled and written.
  the case when the buffer is full and the write goroutine is blocked until the read goroutine is full.
  until the reading bit performs a read. So what it comes down to is
  the proper use of a buffered channel?
  While it is difficult to define it precisely, in one sentence it can be
  expressed as follows: buffered channels should be used when you know the number of running goroutines and you want to limit the number of
  of goroutines that will still be running, or you want to limit the amount of work
  queued for execution.
  Buffered channels work well when you want to either collect data
  from some set of running goroutines, or to limit competitive usage. They can also be used to manage the amount of work
  queued for execution by the system to prevent your services from degrading
  performance and overload your services.

We know exactly how many goroutines are running, and we need each one
goroutine to close when it's finished. This means that we can
create a buffered channel containing one cell for each running goroutine, and allow each goroutine to write its data Ato this
channel without blocking

# When to use atomic operations

Use atomic operations when:

- Simple operations: You need to perform simple operations on a single variable, such as incrementing, decrementing, loading, or storing a value.
- High performance: You want to minimize synchronization overhead for these operations because atomic operations are typically faster than mutexes.
- Secure access: You want secure parallel access to a variable from multiple goroutines.

And something like order count, counting, booleans, small data

#### Examples where atomic operations are appropriate:

1. Counters (increment/decrement)
2. Flags or states (load/save)
3. Safe read/write operations on a single variable

### When to use mutexes

Use mutexes when:

- Complex operations: You need to protect more complex critical sections of code that involve multiple operations or access to multiple variables.
- Multiple variables: You need to synchronize access to multiple variables simultaneously.
- Non-atomic operations: You need to perform operations that cannot be performed atomically.

#### Examples when mutexes are appropriate:

1. Processing structures or objects
2. Updating multiple variables
3. Performing complex operations that require mutual exclusion

### When to choose to use context and sync.WaitGroup?

- Context: Use context to control the execution time of operations, pass values, and cancel operations. This is especially useful in asynchronous operations or when you want to control execution time.

- WaitGroup: sync.WaitGroup is used to wait for multiple goroutines to complete. This is a good way to synchronize the execution of goroutines when you need to wait for a certain number of concurrent tasks to complete.

### benefits of goroutine

- Creating a goroutine takes less time than creating a thread because no system resource is created.
  The initial size of the goroutine stack is smaller than the size of the thread stack and can be
  increased as needed. This makes goroutines more efficient
  in terms of memory utilization.
- Switching between goroutines takes less time than switching between threads because it is done entirely within the process, eliminating the need to make relatively slow
  system calls.
- The scheduler can optimize its decisions because it is an
  an integral part of the Go process. By interacting with the network interrogator,
  scheduler identifies instances where the execution of the go-process should be canceled so that it does not block I/O. It also interacts
  with the garbage collector and makes sure that work is evenly distributed among the operating system threads allocated to your Go process.

### Go language developers support a number of utilities that complement the capabilities of the standard library, collectively referred to as packages

golang.org/x. Among other things, they contain an ErrGroup type created on the basis of the WaitGroup type to get a group of goroutines that terminate processing.
based on the WaitGroup type to produce a group of goroutines that stop processing when one of them returns an error. More detailed
information about this type can be found in the documentation

# When use channel and when mutex

- If you need to coordinate goroutines or track a value across
  its transformation through multiple goroutines, use channels.
- If you need to share a structure field, use mutexes.
- If you have identified a critical performance issue when using channels
  and cannot to find other ways to solve this problem, use mutexes instead of channels mutexes
