## CSP(Communication Sequential Processes)
- Concurrency in Go is based on the paper CSP(Communication Sequential Processes) written by Tony Hoare in 1978.
- CSP is based on 3 core ideas:
    1. Each process is build for sequential execution. Every process has a local state and the process operates on that local state.
    2. If we have to transfer data, data is communicated between processes. No shared memory. We communicate the data, we send a copy of the data over to other processes. Since there is no sharing of memory there will be no Data Races or deadlocks
    3. We can scale easily as each process can run independently. Scale by adding more of the same.
---

## Go's Concurrency Tool Set
- Go-routines:
  - Concurrently executing functions.
- Channels:
  - Are used to communicate data between go-routines.
- Select
  - Is used to multiplex the channel.
  - Multiplex Definitions:
    1. Multiplexing is a technique used to combine and send multiple data streams over a single medium. The process of combining the data streams is known as multiplexing and hardware used for multiplexing is known as a multiplexer.
    2. Multiplexing (sometimes contracted to muxing) is a method by which multiple analog or digital signals are combined into one signal over a shared medium.
    3. Multiplexing is a term used in computer networking, signal processing and telecommunications to refer to the way in which more than one digital or analog signal is combined into just one signal.
    4. The doing of multiple things at the same time or interleaving of many things.
- Sync package:
  - Provides classical synchronization tools(like mutex).
---

## Go-routines
- We can think of go-routines as user space threads managed by the go runtime.
- The go runtime is apart of the executable, it is built into the executable of the application.
- Go-routines are extremely lightweight. Their stacks start with at 2KB, which can shrink and grow as required.
- Go-routines have a low CPU overhead - it takes three instructions per function call.
- We can create hundreds of thousands of go-routines in the same address space.
- Channels are used for communication of data between go-routines. Sharing of memory can be avoided.
- Context switching between go-routines is much cheaper than thread context switching. It happens in the user space.
- The go runtime can be more selective of what data is persisted for retrieval, how it is persisted, and when the persisting needs to occur.
- Go-routine flow:
    1. The Go runtime creates worker OS threads.
    2. Go-routines runs in the context of OS threads. Many go-routines can operate in the context of a single OS thread.
    3. The OS schedules OS threads
    4. The go runtime schedules multiple go-routines on an OS thread. Nothing changes for the OS.