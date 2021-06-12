package select_ex

import (
	"fmt"
	"time"
)

func SecondSelect() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// implement timeout for recv on channel ch
	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
