## Race Condition
- Race condition occurs when order of execution is not guranteed.
- Concurrent programs do not execute in the order they are coded.
---

## WaitGroup
- Go follows the logical concurrency model called 'fork and join'
  - The 'go' statement forks a go-routines(pulls it off of the main routine), when the go-routine is done with its operation is joins back to the main routine.
  - If the main go-routine finishes before the forked go-routine can finish then the program will exit before the go-routine can finish its job.
- WaitGroups allow for us to deterministically block the main go-routine. It creates a join point, and the main go-routine will wait for the forked go-routine to finish before continuing.