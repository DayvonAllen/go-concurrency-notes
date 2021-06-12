## Operating System
- The job of an operating system is to give a fair chance for all processes to access the CPU, memory and other resources.
- What is a process?
  - It's an instance of a running program.
  - It provides an environment for a program to execute.
  - When a program is executed the OS creates a process and allocates memory in a virtual address space.
  - The Virtual address space will contain a Code segment which is compiled machine code(machine instructions)
  - There's a data segment which contains global data.
  - A Heap segment for dynamic memory allocation.
  - A stack segment which is used for storing local variables of functions.
- Process components:
  - Stack
  - Heap
  - Data
  - Code
---

## Misc
- In computing, preemption is the act of temporarily interrupting an executing task, with the intention of resuming it at a later time.
- When a thread gets preempted, it gets interrupted and resumes execution later. It's like a temporary pause before it finishes execution.

## Threads
- Threads are the smallest unit of execution that a CPU accepts.
- Each process has at least on thread(the main thread)
- A process can have multiple threads
- Each thread shares the same address space.
- Each thread has its own stack.
- Threads run independent of each other
- The OS scheduler makes scheduling decisions at the thread level, not at the process level.
- Threads can run concurrently(each thread takes turns on an individual core) or in parallel(Each thread runs at the same time on different cores)
---

## Thread States
- When a process is created the main thread is put into the ready queue(runnable state), it is in the runnable state. 
- Once the CPU is available the thread starts to execute(executing state) and each thread is given a time slice, if that time slice expires, then that thread gets preempted and placed back into the ready queue(runnable state). 
- If the thread gets blocked due to an I/O operation or waiting for an event from other processes, then it's placed into the waiting queue(waiting state) until the I/O operation is complete
- Once it is complete then the thread is placed back into the ready queue(runnable state).
- There are 3 states:
  1. Runnable
  2. Executing
  3. Waiting
- Threads consist of:
  - Stack
  - CPU registers
  - PC(Program counters)
---

## Context Switch
- Process Context consist of:
  - Process State
  - CPU scheduling information
  - Memory Management information
  - Accounting information
  - I/O Status information
- Thread Context consist of:
  - Program counter
  - CPU registers
  - Stack
- Context switches are expensive.
- The CPU has to spend time copying the context of the current executing thread into memory and restoring the context of the next chosen thread.
- It takes thousands of CPU instructions to do context switching and it is a waste of time because when there is context switching going on the CPU is not running your application.
- Context switching of threads in the same process is relatively cheap compared to context switching of threads that are related to a different process.
---

## C10k Problem
- Occurs when we scale the number of threads in a process too high.
- The Scheduler allocates a process a time slice for execution on a CPU core.
- This CPU time slice is divided equally among threads in the process.
    | Scheduler Period | Number of threads | Thread time slice |
    | :--------------- | :---------------: | ----------------- |
    | 10 milliseconds  |         2         | 5 milliseconds    |
    | 10 milliseconds  |         5         | 2 milliseconds    |
    | 10 milliseconds  |       1000        | 10 microseconds   |
- In the last example in the table with 1000 threads, the CPU will spend more time context switching then running the application.
- To do any meaningful job, a thread needs a time slice of at least 2 milliseconds:
    | Scheduler Period | Number of threads | Thread time slice |
    | :--------------- | :---------------: | ----------------- |
    | 2 seconds        |       1000        | 2 milliseconds    |
    | 20 seconds       |      10,000       | 2 milliseconds    |
- If there 1,000 threads it will take 2 seconds to complete 1 cycle(we will have to wait for 2 seconds for the next execution of threads), if there are 10,000 threads then it will take 20 seconds to complete 1 cycle(we will have to wait for 20 seconds for the next execution of threads).
  - In those cases the application will become less responsive.
- Another issue is the stack size(typically it's 8MB), The OS gives a fixed stack size for each thread. The actual size depends on the hardware.
- If you have 8GB in memory(RAM) and the fixed stack size for each thread is 8MB then you can only create 1000 threads.
- The fixed stack size limits the number of threads that we can create to the amount of memory(RAM) that we have.
- In summary:
  - C10K problem = as we scale up the number of threads, the scheduler cycle increases and the application can become less responsive.
