package channels

import "fmt"

func RangeChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			// send iterator over channel
			ch <- i
		}
		close(ch)
	}()

	// range over channel to receive values
	// once the channel is closed the range will exit
	for v := range ch {
		fmt.Println(v)
	}
}
