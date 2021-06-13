package pipeline

import "fmt"

func FirstPipeline() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed
	//ch := generator(2, 3)
	//out := square(ch)

	for v := range square(square(generator(2,3))) {
		fmt.Println(v)
	}
}

// Build a Pipeline
// generator() -> square() -> print

// generator - converts a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <-v * v
		}
		close(out)
	}()
	return out
}
