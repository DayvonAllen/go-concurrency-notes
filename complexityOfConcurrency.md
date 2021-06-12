## Shared Memory
- Threads communicate between each other by sharing memory.
- Sharing of memory between threads creates a lot of complexity.
- Concurrent access to shared memory by two or more threads can lead to a 'Data Race' and the outcome can be 'Un-deterministic'
    - Usually when a thread is writing to a resource while another thread is reading from that same resource
    - The value and behavior can become unpredictable.
---

## Memory Access Synchronization
- It is used to solve the Data Race problem.
- We need to guard the access to shared memory so that a thread gets exclusive access at a time.
- We force the threads to run sequentially when access a particular value or resource.
- We can do this by locking a resource when one thread has it and unlocking the resource when the thread is finished.
- So the developer acquires the lock and releases the lock when done.
- It is a Developer's convention to `lock()` and `unlock()`.
- If the developer does not follow this convention, then we have no guarantee of exclusive access.
- Problems with this approach:
  - A problem of this approach is that it reduces parallelism, Locks force threads to execute sequentially.
  - Inappropriate use of locks can lead to Deadlocks. Circular waits lead to deadlock. One thread has a lock that another thread needs to finish some work and is requesting a lock on a resource that the other thread has, they both go into the waiting state until the other thread is done with the lock but since both threads need each other threads, they are both in the waiting state and they will never release the lock, this is a deadlock. The application will just hang in that case.
---