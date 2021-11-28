## M:N Scheduler
- The Go scheduler is part of the Go runtime.
- It is known as 'M:N Scheduler'
- The Go scheduler runs in the user space.
- It uses OS threads to schedule go-routines for execution.
- Go-routines runs in the context of an OS thread.
- The Go runtime creates a number of worker OS threads, equal to `GOMAXPROCS`.
- `GOMAXPROCS` - default value is the number of processors on your machine.
- The Go scheduler distributes runnable go-routines over multiple worker OS threads.
- At any time, `N` go-routines can be scheduled on `M` OS threads that runs on at most `GOMAXPROCS` number of processors.
---

## Asynchronous Preemption
- As of Go 1.14, The Go scheduler implements asynchronous preemption.
- This prevents long-running go-routines from hogging the CPU, which could block other go-routines.
- The asynchronous preemption is triggered based on a time condition
  - When a go-routine is running for more than 10ms, Go will try to preempt it.
---

## Go-routine State
1. `Runnable` State:
   - When it is created it will be placed in the ready queue and placed in the runnable state.
2. `Executing` state:
   - It moves to the executing state once the go-routine is scheduled on the OS thread.
   - If the go-routine runs through its time slice, it is preempted and placed back into the ready queue.
3. `Waiting` State:
   - If the go-routine is executing, and it gets block on any condition, like blocked on the channel, blocked on a syscall or waiting for a mutex lock, then it is moved to the 'Waiting state'
   - Once the I/O operation is complete it is moved back the Runnable state.
---