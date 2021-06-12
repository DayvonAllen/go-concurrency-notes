package channels

import "fmt"

func BufferedChannel() {
	// buffered channels are non-blocking until the buffer capacity is filled up
	ch := make(chan int, 6)

	go func() {
		defer close(ch)

		// send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
