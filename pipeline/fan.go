package pipeline

import (
	"fmt"
	"sync"
)

func generator2(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
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

func Fan() {
	in := generator2(2, 3)

	// fan out square stage to run two instances.
	ch1 := square2(in)
	ch2 := square2(in)
	// fan in the results of square stages.
	for v := range merge(ch1, ch2) {
		fmt.Println(v)
	}

}

