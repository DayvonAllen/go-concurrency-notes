package cancel

import (
	"fmt"
	"sync"
)

func generator2(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}

		}
	}()
	return out
}

func square2(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}

		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()

		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func CancelGoRou() {
	done := make(chan struct{})
	in := generator2(done, 2, 3)

	// fan out square stage to run two instances.
	ch1 := square2(done, in)
	ch2 := square2(done, in)

	out := merge(done, ch1, ch2)

	// cancel go-routines
	fmt.Println(<-out)
	close(done)
	return
}
